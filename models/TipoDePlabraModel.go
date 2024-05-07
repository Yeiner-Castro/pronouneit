package models

import (
	"gorm.io/gorm"
)

type TipoDePalabra struct {
	gorm.Model
	Grupo string `gorm:"not null"`
}
