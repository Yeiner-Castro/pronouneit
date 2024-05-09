package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func NivelRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/nivel/crear", controllers.NivelCreate)
	routes.GET("/nivel/listar", controllers.NivelGetAll)
	routes.GET("/nivel/buscar/:id", controllers.NivelGetById)
	routes.PUT("/nivel/actualizar/:id", controllers.NivelUpdate)
	routes.DELETE("/nivel/eliminar/:id", controllers.NivelDelete)
	routes.GET("/nivel/maximo", controllers.MaximoNivelGet)

}
