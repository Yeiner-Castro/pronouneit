package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func HistoricoResultadoRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_historico", controllers.HistoricoResultadoCreate)
	routes.GET("/historicos", controllers.HistoricoResultadoGet)
	routes.GET("/buscar_historico_id/:id", controllers.HistoricoResultadoGetById)
	routes.PUT("/actualizar_historico_id/:id", controllers.HistoricoResultadoUpdate)
	routes.DELETE("/elimnar_historico_id/:id", controllers.HistoricoResultadoDelete)

}
