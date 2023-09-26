package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thalisonsilva/gopportun.git/router"
	"net/http"
	
)

func main() {
	// Inicialize Router

	router.Inicializa()
	fmt.Println("Teste de todas")
	
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
}
