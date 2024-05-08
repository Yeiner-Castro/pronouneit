package controllers

import (
	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"time"

	"github.com/gin-gonic/gin"
)

type HistoricoResultadoRequestBody struct {
	Fecha                time.Time `json:"fecha"`
	EjercicioRealizadoID uint      `json:"ejercicioRealizado"`
}

func HistoricoResultadoCreate(c *gin.Context) {
	body := HistoricoResultadoRequestBody{}
	c.BindJSON(&body)

	historicoResultado := &models.HistoricoResultado{
		Fecha:                body.Fecha,
		EjercicioRealizadoID: body.EjercicioRealizadoID,
	}
	result := configs.DB.Create(&historicoResultado)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &historicoResultado)
}

func HistoricoResultadoGet(c *gin.Context) {
	var historicosResultados []models.HistoricoResultado
	configs.DB.Find(&historicosResultados)
	c.JSON(200, &historicosResultados)
	return
}

func HistoricoResultadoGetById(c *gin.Context) {
	id := c.Param("id")
	var historicoResultado models.HistoricoResultado
	configs.DB.First(&historicoResultado, id)
	if historicoResultado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such historical result found"})
		return
	}
	c.JSON(200, &historicoResultado)
	return
}

func HistoricoResultadoUpdate(c *gin.Context) {
	id := c.Param("id")
	var historicoResultado models.HistoricoResultado
	configs.DB.First(&historicoResultado, id)

	if historicoResultado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such historical result found"})
		return
	}

	body := HistoricoResultadoRequestBody{}
	c.BindJSON(&body)
	data := &models.HistoricoResultado{
		Fecha:                body.Fecha,
		EjercicioRealizadoID: body.EjercicioRealizadoID,
	}

	result := configs.DB.Model(&historicoResultado).Updates(data)
	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(200, &historicoResultado)
}

func HistoricoResultadoDelete(c *gin.Context) {
	id := c.Param("id")
	var historicoResultado models.HistoricoResultado
	configs.DB.Delete(&historicoResultado, id)
	if historicoResultado.ID == 0 {
		c.JSON(404, gin.H{"Error": "No such historical result found"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
	return
}
