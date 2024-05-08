package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type EjercicioRealizadoRequestBody struct {
	Resultado   int  `json:"resultado"`
	UsuarioID   uint `json:"usuario"`
	EjercicioID uint `json:"ejercicio"`
}

func EjercicioRealizadoCreate(c *gin.Context) {
	body := EjercicioRealizadoRequestBody{}
	c.BindJSON(&body)

	ejercicioRealizado := &models.EjercicioRealizado{
		Resultado:   body.Resultado,
		UsuarioID:   body.UsuarioID,
		EjercicioID: body.EjercicioID,
	}
	result := configs.DB.Create(&ejercicioRealizado)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &ejercicioRealizado)
}

func EjercicioRealizadoGet(c *gin.Context) {
	var ejerciciosRealizados []models.EjercicioRealizado
	configs.DB.Find(&ejerciciosRealizados)
	c.JSON(200, &ejerciciosRealizados)
	return
}

func EjercicioRealizadoGetById(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	configs.DB.First(&ejercicioRealizado, id)
	if ejercicioRealizado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such completed exercise"})
		return
	}
	c.JSON(200, &ejercicioRealizado)
	return
}

func EjercicioRealizadoUpdate(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	configs.DB.First(&ejercicioRealizado, id)

	if ejercicioRealizado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such completed exercise"})
		return
	}

	body := EjercicioRealizadoRequestBody{}
	c.BindJSON(&body)
	data := &models.EjercicioRealizado{
		Resultado:   body.Resultado,
		UsuarioID:   body.UsuarioID,
		EjercicioID: body.EjercicioID,
	}

	result := configs.DB.Model(&ejercicioRealizado).Updates(data)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &ejercicioRealizado)
}

func EjercicioRealizadoDelete(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	configs.DB.Delete(&ejercicioRealizado, id)
	if ejercicioRealizado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such completed exercise"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
	return
}
