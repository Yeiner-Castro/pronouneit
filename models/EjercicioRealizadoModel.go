package models

import (
	"gorm.io/gorm"
)

type EjercicioRealizado struct {
	gorm.Model
	Resultado   int       `gorm:"not null"`
	UsuarioID   uint      `gorm:"not null"`
	EjercicioID uint      `gorm:"not null"`
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	Ejercicio   Ejercicio `gorm:"foreignKey:EjercicioID"`
}
