package main

import (
	"fmt"

	"example.com/ecommerce/internal/routers"
)

func main() {
	fmt.Println("Start here")
	r := routers.NewRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// func Pong(c *gin.Context) {
// 	// name := c.Param("name")
// 	name := c.DefaultQuery("name", "anonystick")
// 	uid := c.Query("uid")
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "pong....ping " + name,
// 		"uid":     uid,
// 		"users":   []string{"cr7", "m10", "nam"},
// 	})
// }
