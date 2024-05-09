package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func EjercicioRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/ejercicio/crear", controllers.EjercicioCreate)
	routes.GET("/ejercicio/listar", controllers.EjercicioGetAll)
	routes.GET("/ejercicio/buscar/:id", controllers.EjercicioGetById)
	routes.PUT("/ejercicio/actualizar/:id", controllers.EjercicioUpdate)
	routes.DELETE("/ejercicio/eliminar/:id", controllers.EjercicioDelete)

}
