package main

import (
	"log"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	err := configs.DB.AutoMigrate(
		&models.Usuario{},
		&models.EjercicioRealizado{},
		&models.Ejercicio{},
		&models.Nivel{},
		&models.TipoDePalabra{},
	)
	if err != nil {
		log.Fatalf("Error during migration: %s", err)
	}
	log.Println("Migration completed successfully.")
}
