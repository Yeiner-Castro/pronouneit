package controllers

import (
	"net/http"

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

	usuario := &models.Usuario{
		Nombre:      body.Nombre,
		Apellido:    body.Apellido,
		Correo:      body.Correo,
		Contrasenia: body.Contrasenia,
	}

	result := configs.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, &usuario)
}

func UsuarioGetAll(c *gin.Context) {
	var usuarios []models.Usuario
	configs.DB.Find(&usuarios)
	c.JSON(http.StatusOK, &usuarios)
}

func UsuarioGetById(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	result := configs.DB.First(&usuario, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, &usuario)
}

func UsuarioUpdate(c *gin.Context) {

	id := c.Param("id")
	var usuario models.Usuario
	if err := configs.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario not found"})
		return
	}

	var body UsuarioRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updatedData := models.Usuario{
		Nombre:      body.Nombre,
		Apellido:    body.Apellido,
		Correo:      body.Correo,
		Contrasenia: body.Contrasenia,
	}

	result := configs.DB.Model(&usuario).Updates(updatedData)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &usuario)
}

func UsuarioDelete(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	result := configs.DB.Delete(&usuario, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to delete"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
}
