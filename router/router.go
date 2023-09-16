package router

 import "github.com/gin-gonic/gin"

func Inicializa () {
	// Inicializa o Router utilizando as configurações Default de gin
	router := gin.Default()

	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Estamos rodando a nossa api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}