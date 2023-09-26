package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thalisonsilva/gopportun.git/router"
	
)

func main() {

	// Inicialize Router
	router.Inicializa()
}
func setupRouter(*gin.Engine){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}