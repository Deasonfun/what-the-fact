package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client := &http.Client{}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"fact": GetFact(*client),
		})
	})
	r.Run()
}

func GetFact(client http.Client) (fact string) {
	req, _ := http.NewRequest("GET", "https://api.api-ninjas.com/v1/facts?limit=1", nil)
	req.Header.Set("X-Api-Key", "6Ioq8wPk4a8CSiaBhn7lUw==W3UHGZPwOGyCT5z2")
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	reqData := string(body)
	data := strings.Split(reqData, ":")
	factData := strings.Split(data[1], ":")
	factString := strings.Split(factData[0], "}")
	fmt.Println(strings.TrimLeft(factString[0], "\""))
	fact = strings.TrimLeft(factString[0], "\"")
	return
}
