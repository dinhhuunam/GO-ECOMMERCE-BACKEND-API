package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start here")
	r := gin.Default()

	v1 := r.Group("/v1/")
	{
		v1.GET("/ping/:name", Pong)
		v1.PUT("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)
	}

	v2 := r.Group("/v2/")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Pong(c *gin.Context) {
	// name := c.Param("name")
	name := c.DefaultQuery("name", "anonystick")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong....ping " + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "nam"},
	})
}
