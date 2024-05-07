package models

import (
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre      string `gorm:"size:50;not null"`
	Apellido    string `gorm:"size:50;not null"`
	Correo      string `gorm:"size:50;not null;unique"`
	Contrasenia string `gorm:"size:50;not null"`
}
