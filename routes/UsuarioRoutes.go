package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/usuario/crear", controllers.UsuarioCreate)
	routes.GET("/usuario/listar", controllers.UsuarioGetAll)
	routes.GET("/usuario/buscar/:id", controllers.UsuarioGetById)
	routes.PUT("/usuario/actualizar/:id", controllers.UsuarioUpdate)
	routes.DELETE("/usuario/eliminar/:id", controllers.UsuarioDelete)
	router.GET("/usuario/:usuarioId/nivelActual", controllers.GetUsuarioNivelActual)

}
