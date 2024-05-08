package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func EjercicioRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_ejercicio", controllers.EjercicioCreate)
	routes.GET("/ejercicios", controllers.EjercicioGetAll)
	routes.GET("/buscar_ejercicio_id/:id", controllers.EjercicioGetById)
	routes.PUT("/actualizar_ejercicio_id/:id", controllers.EjercicioUpdate)
	routes.DELETE("/elimnar_ejercicio_id/:id", controllers.EjercicioDelete)

}
