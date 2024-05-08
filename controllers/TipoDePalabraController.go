package controllers

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type TipoDePalabraRequestBody struct {
	Grupo string `json:"grupo"`
}

func TipoDePalabraCreate(c *gin.Context) {
	var body TipoDePalabraRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	tipoDePalabra := models.TipoDePalabra{
		Grupo: body.Grupo,
	}

	result := configs.DB.Create(&tipoDePalabra)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &tipoDePalabra)
}

func TipoDePalabraGetAll(c *gin.Context) {
	var tiposDePalabra []models.TipoDePalabra
	configs.DB.Find(&tiposDePalabra)
	c.JSON(http.StatusOK, &tiposDePalabra)
}

func TipoDePalabraGetById(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	result := configs.DB.First(&tipoDePalabra, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such tipo de palabra found"})
		return
	}

	c.JSON(http.StatusOK, &tipoDePalabra)
}

func TipoDePalabraUpdate(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	configs.DB.First(&tipoDePalabra, id)

	var body TipoDePalabraRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	data := &models.TipoDePalabra{
		Grupo: body.Grupo,
	}

	result := configs.DB.Model(&tipoDePalabra).Updates(data)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, &tipoDePalabra)
}

func TipoDePalabraDelete(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	result := configs.DB.Delete(&tipoDePalabra, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
