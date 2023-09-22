package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thalisonsilva/gopportun.git/router"
)

func main() {
	// Inicialize Router

	router.Inicializa()	
	router.InicializaRoutes()
}
