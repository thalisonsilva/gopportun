package router

 import "github.com/gin-gonic/gin"

func Inicializa () {

	// Inicialize Routes
	router := gin.Default()

	// Inicialize Routes
	
	// Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}