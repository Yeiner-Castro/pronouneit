package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func EjercicioRealizadoRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_ejercicio_realizado", controllers.EjercicioRealizadoCreate)
	routes.GET("/ejercicios_realizados", controllers.EjercicioRealizadoGetAll)
	routes.GET("/buscar_ejercicio_realizado_id/:id", controllers.EjercicioRealizadoGetById)
	routes.PUT("/actualizar_ejercicio_realizado_id/:id", controllers.EjercicioRealizadoUpdate)
	routes.DELETE("/elimnar_ejercicio_realizado_id/:id", controllers.EjercicioRealizadoDelete)
}
