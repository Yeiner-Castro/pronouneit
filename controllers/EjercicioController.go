package controllers

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type EjercicioRequestBody struct {
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
	NivelID   uint   `json:"nivelId"`
	TipoID    uint   `json:"tipoId"`
}

func EjercicioCreate(c *gin.Context) {
	body := EjercicioRequestBody{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid body"})
		return
	}

	ejercicio := &models.Ejercicio{
		Nombre:    body.Nombre,
		Contenido: body.Contenido,
		NivelID:   body.NivelID,
		TipoID:    body.TipoID,
	}
	result := configs.DB.Create(&ejercicio)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(http.StatusOK, &ejercicio)
}

func EjercicioGetAll(c *gin.Context) {
	var ejercicios []models.Ejercicio
	configs.DB.Preload("Nivel").Preload("Tipo").Find(&ejercicios)
	c.JSON(http.StatusOK, &ejercicios)
}

func EjercicioGetById(c *gin.Context) {
	id := c.Param("id")
	var ejercicio models.Ejercicio
	result := configs.DB.Preload("Nivel").Preload("Tipo").First(&ejercicio, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such exercise"})
		return
	}

	c.JSON(http.StatusOK, &ejercicio)
}

func EjercicioUpdate(c *gin.Context) {
	id := c.Param("id")
	var ejercicio models.Ejercicio
	err := configs.DB.First(&ejercicio, id)
	if err.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "No such exercise"})
		return
	}

	var body EjercicioRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updateData := &models.Ejercicio{
		Nombre:    body.Nombre,
		Contenido: body.Contenido,
		NivelID:   body.NivelID,
		TipoID:    body.TipoID,
	}

	result := configs.DB.Model(&ejercicio).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &ejercicio)
}

func EjercicioDelete(c *gin.Context) {
	id := c.Param("id")
	result := configs.DB.Delete(&models.Ejercicio{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ejercicio"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func EjercicioGetByNivel(c *gin.Context) {
	nivelId := c.Param("nivelId")
	var ejercicios []models.Ejercicio
	configs.DB.Preload("Nivel").Preload("Tipo").Where("nivel_id = ?", nivelId).Find(&ejercicios)
	c.JSON(http.StatusOK, &ejercicios)
}
