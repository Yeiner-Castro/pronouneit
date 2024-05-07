package models

import (
	"time"

	"gorm.io/gorm"
)

type HistoricoResultado struct {
	gorm.Model
	Fecha                time.Time          `gorm:"not null"`
	EjercicioRealizadoID uint               `gorm:"not null"`
	EjercicioRealizado   EjercicioRealizado `gorm:"foreignKey:EjercicioRealizadoID"`
}
