package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func TipoDePalabraRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/tipo_de_palabra/crear", controllers.TipoDePalabraCreate)
	routes.GET("/tipo_de_palabra/listar", controllers.TipoDePalabraGetAll)
	routes.GET("/tipo_de_palabra/buscar/:id", controllers.TipoDePalabraGetById)
	routes.PUT("/tipo_de_palabra/actualizar/:id", controllers.TipoDePalabraUpdate)
	routes.DELETE("/tipo_de_palabra/eliminar/:id", controllers.TipoDePalabraDelete)

}
