package main

import (
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	firstName := c.DefaultQuery("firstName", "Guest")
	lastName := c.Query("lastName")
	c.String(200, "Hello %s %s", firstName, lastName)

}

func main() {
	router := gin.Default()
	router.GET("/hello", HelloHandler)
	router.Run(":8080")
}
