package models

import (
	"gorm.io/gorm"
)

type EjercicioRealizado struct {
	gorm.Model
	Resultado   int       `gorm:"not null"`
	UsuarioID   uint      `gorm:"not null"`
	EjercicioID uint      `gorm:"not null"`
	Aprobado    bool      `gorm:"default:false"` // Añade esto
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID"`
	Ejercicio   Ejercicio `gorm:"foreignKey:EjercicioID"`
}
