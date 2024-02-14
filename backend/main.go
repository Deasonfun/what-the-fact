package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"math/rand"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//Load .env file
	err := godotenv.Load()
	CheckError(err)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	//Set up string that will be used to connect to SQL
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Start backend web server
	router := gin.Default()

	//Set up CORS to accept frontend queries only from localhost:3000
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
	}))

	//Create a API route /getfact
	router.GET("/getfact", func(c *gin.Context) {
		//When status is OK, send out a package with the fact
		c.JSON(http.StatusOK, gin.H{
			//Run the GetFact function and send out the return
			"fact": GetFact(connectString),
		})
	})

	//Run the web server on port 8080
	router.Run(":8080")
}

// GetFact funtion will query the SQL database for a fact
// Accepts the connectionString (string) and returns a fact (string)
func GetFact(connectString string) (fact string) {
	//Open the database
	db, err := sql.Open("postgres", connectString)
	CheckError(err)

	var count int

	//Query the db to get a count of all the rows
	err = db.QueryRow("SELECT COUNT(*) FROM facts").Scan(&count)
	CheckError(err)

	//Get the random ID
	var randId = rand.Intn(count - 1) + 1

	var queryFact string
	//Query the database for a fact using the random ID
	var query = fmt.Sprintf("SELECT fact FROM facts WHERE id=%d", randId)
	rows, err := db.Query(query)
	CheckError(err)
	defer rows.Close()
	//Scan the query and set it to the the queryFact variable
	for rows.Next() {
		err = rows.Scan(&queryFact)
		CheckError(err)
	}
	//Check the query for a SQL error
	err = rows.Err()
	CheckError(err)
	//Return fact from database
	return queryFact
}

// All errors will be sent to this function to be checked
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
