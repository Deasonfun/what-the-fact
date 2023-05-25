package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	CheckError(err)
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	BuildDB(connectString)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
	}))

	client := &http.Client{}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"fact": GetFact(*client, connectString),
		})
	})

	r.Run(":8080")
}

func BuildDB(connectString string) {

	db, err := sql.Open("postgres", connectString)
	CheckError(err)

	defer db.Close()
	err = db.Ping()
	CheckError(err)
	fmt.Println("Connected to database...")

	factJSON, err := ioutil.ReadFile("./facts.json")
	CheckError(err)

	var payload map[string]interface{}
	err = json.Unmarshal([]byte(factJSON), &payload)
	CheckError(err)

	fmt.Println(payload["fact"])
	fmt.Println("")

	_, err = db.Exec(fmt.Sprintf("INSERT INTO facts (fact) VALUES('%s')", payload["fact"]))
	CheckError(err)
}

func GetFact(client http.Client, connectString string) (fact string) {
	var (
		queryFact string
	)
	db, err := sql.Open("postgres", connectString)
	CheckError(err)
	rows, err := db.Query("select fact from facts where id=1")
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&queryFact)
		CheckError(err)
	}
	err = rows.Err()
	CheckError(err)
	return queryFact
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
