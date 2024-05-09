package controllers

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type NivelRequestBody struct {
	Nivel int `json:"nivel"`
}

func NivelCreate(c *gin.Context) {
	var body NivelRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	nivel := models.Nivel{
		Nivel: body.Nivel,
	}

	result := configs.DB.Create(&nivel)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, &nivel)
}

func NivelGetAll(c *gin.Context) {
	var niveles []models.Nivel
	configs.DB.Find(&niveles)
	c.JSON(http.StatusOK, &niveles)
}

func NivelGetById(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	result := configs.DB.First(&nivel, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"Error": "No such nivel found"})
		return
	}

	c.JSON(200, &nivel)
}

func NivelUpdate(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	err := configs.DB.First(&nivel, id)
	if err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such nivel found"})
		return
	}

	var body NivelRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updatedData := &models.Nivel{
		Nivel: body.Nivel,
	}

	result := configs.DB.Model(&nivel).Updates(updatedData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, &nivel)
}

func NivelDelete(c *gin.Context) {
	id := c.Param("id")
	var nivel models.Nivel
	result := configs.DB.Delete(&nivel, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No such nivel found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func MaximoNivelGet(c *gin.Context) {
	var maxNivel models.Nivel
	result := configs.DB.Order("nivel desc").First(&maxNivel)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "No se pudo encontrar el nivel más alto"})
		return
	}

	c.JSON(200, gin.H{"max_nivel": maxNivel.Nivel})
}
