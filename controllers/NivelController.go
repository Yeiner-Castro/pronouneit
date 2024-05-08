package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type NivelRequestBody struct {
	Nivel int `json:"nivel"`
}

func NivelCreate(c *gin.Context) {
	body := NivelRequestBody{}
	c.BindJSON(&body)

	nivel := &models.Nivel{Nivel: body.Nivel}
	result := configs.DB.Create(&nivel)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &nivel)
}

func NivelGet(c *gin.Context) {
	var niveles []models.Nivel
	configs.DB.Find(&niveles)
	c.JSON(200, &niveles)
	return
}

func NivelGetById(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	configs.DB.First(&nivel, id)
	c.JSON(200, &nivel)
	return
}

func NivelUpdate(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	configs.DB.First(&nivel, id)

	body := NivelRequestBody{}
	c.BindJSON(&body)
	data := &models.Nivel{Nivel: body.Nivel}

	result := configs.DB.Model(&nivel).Updates(data)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &nivel)
}

func NivelDelete(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	configs.DB.Delete(&nivel, id)
	c.JSON(200, gin.H{"deleted": true})
	return
}
