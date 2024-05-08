package routes

import (
	"github.com/Yeiner-Castro/pronouneit.git/controllers"

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.Engine) {

	routes := router.Group("/")
	routes.POST("/crear_usuarios", controllers.UsuarioCreate)
	routes.GET("/usuarios", controllers.UsuarioGetAll)
	routes.GET("/buscar_usuario_id/:id", controllers.UsuarioGetById)
	routes.PUT("/actualizar_usuario_id/:id", controllers.UsuarioUpdate)
	routes.DELETE("/elimnar_usuario_id/:id", controllers.UsuarioDelete)

}
