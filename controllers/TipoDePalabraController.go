package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type TipoDePalabraRequestBody struct {
	Grupo string `json:"grupo"`
}

func TipoDePalabraCreate(c *gin.Context) {
	body := TipoDePalabraRequestBody{}
	c.BindJSON(&body)

	tipoDePalabra := &models.TipoDePalabra{Grupo: body.Grupo}
	result := configs.DB.Create(&tipoDePalabra)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &tipoDePalabra)
}

func TipoDePalabraGet(c *gin.Context) {
	var tiposDePalabra []models.TipoDePalabra
	configs.DB.Find(&tiposDePalabra)
	c.JSON(200, &tiposDePalabra)
	return
}

func TipoDePalabraGetById(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	configs.DB.First(&tipoDePalabra, id)
	c.JSON(200, &tipoDePalabra)
	return
}

func TipoDePalabraUpdate(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	configs.DB.First(&tipoDePalabra, id)

	body := TipoDePalabraRequestBody{}
	c.BindJSON(&body)
	data := &models.TipoDePalabra{Grupo: body.Grupo}

	result := configs.DB.Model(&tipoDePalabra).Updates(data)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &tipoDePalabra)
}

func TipoDePalabraDelete(c *gin.Context) {
	id := c.Param("id")
	var tipoDePalabra models.TipoDePalabra
	configs.DB.Delete(&tipoDePalabra, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
