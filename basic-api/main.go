package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
    engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
		    "message": "hello gin",
		})
    })
	engine.Run(":3000")
}
