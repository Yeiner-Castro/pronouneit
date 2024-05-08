package controllers

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type EjercicioRealizadoRequestBody struct {
	Resultado   int  `json:"resultado"`
	UsuarioID   uint `json:"usuarioId"`
	EjercicioID uint `json:"ejercicioId"`
}

func EjercicioRealizadoCreate(c *gin.Context) {
	var body EjercicioRealizadoRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid body request"})
		return
	}

	ejercicioRealizado := &models.EjercicioRealizado{
		Resultado:   body.Resultado,
		UsuarioID:   body.UsuarioID,
		EjercicioID: body.EjercicioID,
	}
	result := configs.DB.Create(&ejercicioRealizado)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, &ejercicioRealizado)
}

func EjercicioRealizadoGetAll(c *gin.Context) {
	var ejerciciosRealizados []models.EjercicioRealizado
	configs.DB.Preload("Usuario").Preload("Ejercicio").Find(&ejerciciosRealizados)
	c.JSON(http.StatusOK, &ejerciciosRealizados)
}

func EjercicioRealizadoGetById(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	result := configs.DB.Preload("Usuario").Preload("Ejercicio").First(&ejercicioRealizado, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such completed exercise"})
		return
	}
	c.JSON(http.StatusOK, &ejercicioRealizado)
}

func EjercicioRealizadoUpdate(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	result := configs.DB.First(&ejercicioRealizado, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such completed exercise"})
		return
	}

	var body EjercicioRealizadoRequestBody
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid body request"})
		return
	}

	updatedData := &models.EjercicioRealizado{
		Resultado:   body.Resultado,
		UsuarioID:   body.UsuarioID,
		EjercicioID: body.EjercicioID,
	}

	result2 := configs.DB.Model(&ejercicioRealizado).Updates(updatedData)
	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, &ejercicioRealizado)
}

func EjercicioRealizadoDelete(c *gin.Context) {
	id := c.Param("id")
	var ejercicioRealizado models.EjercicioRealizado
	result := configs.DB.Delete(&ejercicioRealizado, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "No such completed exercise"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
