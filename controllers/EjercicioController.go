package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
)

type EjercicioRequestBody struct {
	Nombre    string `json:"nombre"`
	Contenido string `json:"contenido"`
	NivelID   uint   `json:"nivel"`
	TipoID    uint   `json:"grupo"`
}

func EjercicioCreate(c *gin.Context) {
	body := EjercicioRequestBody{}
	c.BindJSON(&body)

	ejercicio := &models.Ejercicio{
		Nombre:    body.Nombre,
		Contenido: body.Contenido,
		NivelID:   body.NivelID,
		TipoID:    body.TipoID,
	}
	result := configs.DB.Create(&ejercicio)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &ejercicio)
}

func EjercicioGet(c *gin.Context) {
	var ejercicios []models.Ejercicio
	configs.DB.Find(&ejercicios)
	c.JSON(200, &ejercicios)
	return
}

func EjercicioGetById(c *gin.Context) {
	id := c.Param("id")
	var ejercicio models.Ejercicio
	configs.DB.First(&ejercicio, id)
	if ejercicio.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such exercise"})
		return
	}
	c.JSON(200, &ejercicio)
	return
}

func EjercicioUpdate(c *gin.Context) {
	id := c.Param("id")
	var ejercicio models.Ejercicio
	configs.DB.First(&ejercicio, id)

	if ejercicio.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such exercise"})
		return
	}

	body := EjercicioRequestBody{}
	c.BindJSON(&body)
	data := &models.Ejercicio{
		Nombre:    body.Nombre,
		Contenido: body.Contenido,
		NivelID:   body.NivelID,
		TipoID:    body.TipoID,
	}

	result := configs.DB.Model(&ejercicio).Updates(data)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &ejercicio)
}

func EjercicioDelete(c *gin.Context) {
	id := c.Param("id")
	var ejercicio models.Ejercicio
	configs.DB.Delete(&ejercicio, id)
	if ejercicio.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such exercise"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
	return
}
