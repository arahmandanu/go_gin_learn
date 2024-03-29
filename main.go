package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println(c.Request.URL.Query())
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
