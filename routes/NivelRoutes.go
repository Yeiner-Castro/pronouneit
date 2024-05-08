package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func NivelRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_nivel", controllers.NivelCreate)
	routes.GET("/niveles", controllers.NivelGetAll)
	routes.GET("/buscar_nivel_id/:id", controllers.NivelGetById)
	routes.PUT("/actualizar_nivel_id/:id", controllers.NivelUpdate)
	routes.DELETE("/elimnar_nivel_id/:id", controllers.NivelDelete)

}
