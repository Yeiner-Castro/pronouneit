package controllers

import (
	"net/http"

	"github.com/Yeiner-Castro/pronouneit.git/configs"
	"github.com/Yeiner-Castro/pronouneit.git/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioRequestBody struct {
	Nombre      string `json:"nombre" validate:"required"`
	Apellido    string `json:"apellido" validate:"required"`
	Correo      string `json:"correo" validate:"required,email"`
	Contrasenia string `json:"contrasenia" validate:"required,min=8"`
}

type CambioContraseniaRequest struct {
	ContraseniaActual string `json:"contrasenia_actual" validate:"required,min=8"`
	NuevaContrasenia  string `json:"nueva_contrasenia" validate:"required,min=8"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func UsuarioCreate(c *gin.Context) {
	var body UsuarioRequestBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	hashedPassword, err := hashPassword(body.Contrasenia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	usuario := models.Usuario{
		Nombre:      body.Nombre,
		Apellido:    body.Apellido,
		Correo:      body.Correo,
		Contrasenia: hashedPassword,
	}

	result := configs.DB.Create(&usuario)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to insert", "details": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": usuario.ID, "nombre": usuario.Nombre, "apellido": usuario.Apellido, "correo": usuario.Correo})
}

func UsuarioGetAll(c *gin.Context) {
	var usuarios []models.Usuario
	configs.DB.Find(&usuarios)
	c.JSON(http.StatusOK, &usuarios)
}

func UsuarioGetById(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	result := configs.DB.First(&usuario, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": usuario.ID, "nombre": usuario.Nombre, "apellido": usuario.Apellido, "correo": usuario.Correo})
}

func UsuarioUpdate(c *gin.Context) {

	id := c.Param("id")
	var usuario models.Usuario
	if err := configs.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario not found"})
		return
	}

	var body UsuarioRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	updatedData := models.Usuario{
		Nombre:   body.Nombre,
		Apellido: body.Apellido,
		Correo:   body.Correo,
	}

	result := configs.DB.Model(&usuario).Updates(updatedData)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": true, "message": "Failed to update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": usuario.ID, "nombre": usuario.Nombre, "apellido": usuario.Apellido, "correo": usuario.Correo})
}

func UsuarioDelete(c *gin.Context) {
	id := c.Param("id")
	var usuario models.Usuario
	result := configs.DB.Delete(&usuario, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func GetUsuarioNivelActual(c *gin.Context) {
	userID := c.Param("usuarioId")

	// Primero, obtenemos el nivel máximo de la base de datos
	var maxNivel models.Nivel
	if err := configs.DB.Order("nivel DESC").First(&maxNivel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve maximum level"})
		return
	}

	// Determinar el nivel actual del usuario verificando todos los ejercicios de cada nivel
	nivelActual := 0
	for nivel := 1; nivel <= maxNivel.Nivel; nivel++ {
		var ejercicios []models.Ejercicio
		configs.DB.Where("nivel_id = ?", nivel).Find(&ejercicios)

		todosAprobados := true
		for _, ejercicio := range ejercicios {
			// Verificar si existe al menos un intento aprobado para este ejercicio
			var count int64
			configs.DB.Model(&models.EjercicioRealizado{}).
				Where("usuario_id = ? AND ejercicio_id = ? AND aprobado = ?", userID, ejercicio.ID, true).
				Count(&count)

			if count == 0 {
				todosAprobados = false
				break
			}
		}

		if todosAprobados {
			nivelActual = nivel
		} else {
			nivelActual++
			break // Si no todos los ejercicios de este nivel están aprobados, detiene el bucle
		}
	}

	c.JSON(http.StatusOK, gin.H{"nivelActual": nivelActual})
}

func CambiarContrasenia(c *gin.Context) {
	usuarioID := c.Param("id") // O obtenerlo de la sesión del usuario autenticado
	var body CambioContraseniaRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	var usuario models.Usuario
	if err := configs.DB.First(&usuario, usuarioID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario not found"})
		return
	}

	// Verificar que la contraseña actual sea correcta
	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasenia), []byte(body.ContraseniaActual)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Contraseña actual incorrecta"})
		return
	}

	// Hash de la nueva contraseña
	nuevaContraseniaHashed, err := bcrypt.GenerateFromPassword([]byte(body.NuevaContrasenia), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// Actualizar la contraseña en la base de datos
	usuario.Contrasenia = string(nuevaContraseniaHashed)
	if err := configs.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contraseña actualizada correctamente"})
}
