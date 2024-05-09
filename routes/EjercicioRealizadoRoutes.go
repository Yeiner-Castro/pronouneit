package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func EjercicioRealizadoRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/ejercicio_realizado/crear", controllers.EjercicioRealizadoCreate)
	routes.GET("/ejercicios_realizado/listar", controllers.EjercicioRealizadoGetAll)
	routes.GET("/ejercicio_realizado/buscar/:id", controllers.EjercicioRealizadoGetById)
	routes.PUT("/ejercicio_realizado/actualizar/:id", controllers.EjercicioRealizadoUpdate)
	routes.DELETE("/ejercicio_realizado/eliminar/:id", controllers.EjercicioRealizadoDelete)
	router.GET("/ejercicios_realizados/usuario/:usuarioId/ejercicio/:ejercicioId/ultimo", controllers.GetUltimoEjercicioRealizadoByUsuario)
	router.GET("/ejercicios_realizados/usuario/:usuarioId/detalles", controllers.GetEjerciciosRealizadosDetalleByUsuario)

}
