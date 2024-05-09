package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func TipoDePalabraRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_grupo", controllers.TipoDePalabraCreate)
	routes.GET("/tipo_de_palabras", controllers.TipoDePalabraGetAll)
	routes.GET("/buscar_grupo_id/:id", controllers.TipoDePalabraGetById)
	routes.PUT("/actualizar_grupo_id/:id", controllers.TipoDePalabraUpdate)
	routes.DELETE("/eliminar_grupo_id/:id", controllers.TipoDePalabraDelete)

}
