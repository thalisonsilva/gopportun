package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InicializaRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1") // Correção 1: Usar := para declarar v1 sem usar *gin.RouterGroup
	{	
		
		v1.GET("/opening", func(ctx *gin.Context) { // Correção 2: Colocar a função dentro do GET
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "POST Opening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) { // Correção 2: Colocar a função dentro do GET
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "POST Opening",
			})
		})	
		v1.DELETE("/opening", func(ctx *gin.Context) { // Correção 2: Colocar a função dentro do GET
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "DELETE Opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context) { // Correção 2: Colocar a função dentro do GET
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "PUT Opening",
			})
		})
		v1.GET("/opening", func(ctx *gin.Context) { // Correção 2: Colocar a função dentro do GET
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "GET Openings",
			})
		})
	}
}
