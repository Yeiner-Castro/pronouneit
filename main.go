package main

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()

	routes.UsuarioRouter(r)
	routes.TipoDePalabraRouter(r)
	routes.EjercicioRealizadoRouter(r)
	routes.EjercicioRouter(r)
	routes.NivelRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Goooo.",
		})
	})
	r.Run()
}
