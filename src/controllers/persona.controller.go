package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"modcrudmysql.com/src/commons"
	"modcrudmysql.com/src/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	db := commons.GetConnection()
	//cerrar la conexion al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al obtener la conexión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	defer sqlDB.Close()

	var personas []models.Persona
	if err := db.Find(&personas).Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al recuperar personas: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}

	// Envía un mensaje de éxito junto con el estado 200
	successMessage := "Solicitud completada correctamente."
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
		"data":    personas,
	})

}

func GetById(w http.ResponseWriter, r *http.Request) {
	db := commons.GetConnection()
	//cerrar la conexion al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al obtener la conexión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	defer sqlDB.Close()

	var persona models.Persona
	if err := db.First(&persona, r.URL.Query().Get("id")).Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al recuperar persona: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}

	// Envía un mensaje de éxito junto con el estado 200
	successMessage := "Solicitud completada correctamente."
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
		"data":    persona,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	db := commons.GetConnection()
	// Cerrar la conexión al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al obtener la conexión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	defer sqlDB.Close()

	var persona models.Persona

	// Decodificar el cuerpo JSON en la estructura Persona
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&persona); err != nil {
		errorMessage := "Error al decodificar los datos JSON: " + err.Error()
		commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
		return
	}

	if err := db.Create(&persona).Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al registrar persona: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}

	// Envía un mensaje de éxito junto con el estado 200
	successMessage := "Solicitud completada correctamente."
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
		"data":    persona,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := commons.GetConnection()
	// Cerrar la conexión al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al obtener la conexión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	defer sqlDB.Close()

	var persona models.Persona

	// Decodificar el cuerpo JSON en la estructura Persona
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&persona); err != nil {
		errorMessage := "Error al decodificar los datos JSON: " + err.Error()
		commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
		return
	}

	if err := db.Where("email = ? AND password = ?", persona.Email, persona.Password).First(&persona).Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al iniciar sesión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}

	// Envía un mensaje de éxito junto con el estado 200
	successMessage := "Solicitud completada correctamente."
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
		"data":    persona,
	})
}


func Update(w http.ResponseWriter, r *http.Request) {
    db := commons.GetConnection()
    // Cerrar la conexión al final
    sqlDB, err := db.DB()
    if err != nil {
        // Manejar el error de la base de datos
        errorMessage := "Error al obtener la conexión: " + err.Error()
        commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
        return
    }
    defer sqlDB.Close()

    // Obtener el id del parámetro de ruta
    id, ok := mux.Vars(r)["id"]
    if !ok {
        errorMessage := "el parámetro id es requerido"
        commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
        return
    }

    // Verificar si el registro existe antes de actualizar
    var persona models.Persona
    if err := db.First(&persona, id).Error; err != nil {
        errorMessage := "La persona con ID " + id + " no existe."
        commons.SendError(w, http.StatusNotFound, errors.New(errorMessage))
        return
    }

    // Decodificar el cuerpo JSON en la estructura Persona
    var updatedFields map[string]interface{}
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&updatedFields); err != nil {
        errorMessage := "Error al decodificar los datos JSON: " + err.Error()
        commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
        return
    }

    // Actualizar solo los campos proporcionados en el JSON
    if err := db.Model(&persona).Where("id = ?", id).Updates(updatedFields).Error; err != nil {
        // Manejar el error de la base de datos
        errorMessage := "Error al actualizar persona: " + err.Error()
        commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
        return
    }

    // Envía un mensaje de éxito junto con el estado 200 y devuelve el registro actualizado
    successMessage := "Solicitud completada correctamente"
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
		"data":    persona,
	})
}


func Delete(w http.ResponseWriter, r *http.Request) {
	db := commons.GetConnection()
	//cerrar la conexion al final
	sqlDB, err := db.DB()
	if err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al obtener la conexión: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	defer sqlDB.Close()

	var persona models.Persona
	// Obtener el id del parámetro de ruta
	id, ok := mux.Vars(r)["id"]
	if !ok {
		errorMessage := "el parámetro id es requerido"
		commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
		return
	}
	// Validar que el id sea un número entero válido
	if _, err := strconv.Atoi(id); err != nil {
		errorMessage := "el parámetro id debe ser un número entero válido"
		commons.SendError(w, http.StatusBadRequest, errors.New(errorMessage))
		return
	}
	// Verificar la existencia de la persona en la base de datos
	var count int64
	if err := db.Model(&models.Persona{}).Where("id = ?", id).Count(&count).Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al verificar la existencia de la persona: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}
	if count == 0 {
		// La persona no existe, devolver un error
		errorMessage := fmt.Sprintf("La persona con el id %s no existe.", id)
		commons.SendError(w, http.StatusNotFound, errors.New(errorMessage))
		return
	}
	// Log the delete operation
	log.Printf("Deleting persona with id %s", id)
	// Obtener el resultado de la operación Delete
	result := db.Where("id = ?", id).Delete(&persona)
	if err := result.Error; err != nil {
		// Manejar el error de la base de datos
		errorMessage := "Error al eliminar persona: " + err.Error()
		commons.SendError(w, http.StatusInternalServerError, errors.New(errorMessage))
		return
	}

	// Obtener el número de filas afectadas
	rowsAffected := result.RowsAffected

	// Construir el mensaje de éxito
	successMessage := fmt.Sprintf("Usuario con el id %s ha sido eliminado. Filas afectadas: %d", id, rowsAffected)

	// Envía un mensaje de éxito junto con el estado 200
	commons.SendResponse(w, http.StatusOK, map[string]interface{}{
		"message": successMessage,
	})
}
