package main

import (
	"github.com/rhettli/gin"
	"github.com/thinkerou/favicon"
)

func main() {
	app := gin.Default()
	app.Use(favicon.New("./favicon.ico"))
	app.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello favicon.")
	})
	app.Run(":8080")
}
