package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type UsuarioRequestBody struct {
	Nombre      string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Correo      string `json:"correo"`
	Contrasenia string `json:"contrasenia"`
}

func UsuarioCreate(c *gin.Context) {

	body := UsuarioRequestBody{}

	c.BindJSON(&body)

	usuario := &models.Usuario{Nombre: body.Nombre, Apellido: body.Apellido, Correo: body.Correo, Contrasenia: body.Contrasenia}

	result := configs.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &usuario)
}

func UsuarioGet(c *gin.Context) {
	var usuarios []models.Usuario
	configs.DB.Find(&usuarios)
	c.JSON(200, &usuarios)
	return
}

func UsuarioGetById(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.First(&usuario, id)
	c.JSON(200, &usuario)
	return
}

func UsuarioUpdate(c *gin.Context) {

	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.First(&usuario, id)

	body := UsuarioRequestBody{}
	c.BindJSON(&body)
	data := &models.Usuario{Nombre: body.Nombre, Apellido: body.Apellido, Correo: body.Correo, Contrasenia: body.Contrasenia}

	result := configs.DB.Model(&usuario).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &usuario)
}

func UsuarioDelete(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	configs.DB.Delete(&usuario, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
