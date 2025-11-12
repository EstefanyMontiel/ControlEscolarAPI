package handlers

import (
    "net/http"
    "strconv"
    "log"
    
    "github.com/gin-gonic/gin"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"
)

// CreateStudent maneja POST /api/students
// CreateStudent maneja POST /api/students
func CreateStudent(c *gin.Context) {
    var student models.Student
    
    // Validar el JSON de entrada
    if err := c.ShouldBindJSON(&student); err != nil {
        // Log del error para debugging
        log.Printf("Error en validación: %v", err)
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    // Crear el estudiante en la base de datos
    if err := config.GetDB().Create(&student).Error; err != nil {
        log.Printf("Error al crear estudiante: %v", err)
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear el estudiante")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusCreated, "Estudiante creado exitosamente", student)
}


// GetAllStudents maneja GET /api/students
func GetAllStudents(c *gin.Context) {
    var students []models.Student
    
    if err := config.GetDB().Find(&students).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener estudiantes")
        return
    }
    
    c.JSON(http.StatusOK, students)
}

// GetStudent maneja GET /api/students/:student_id
func GetStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("student_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var student models.Student
    if err := config.GetDB().First(&student, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    c.JSON(http.StatusOK, student)
}

// UpdateStudent maneja PUT /api/students/:student_id
func UpdateStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("student_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var student models.Student
    if err := config.GetDB().First(&student, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    var updatedData models.Student
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    // Actualizar campos
    student.Name = updatedData.Name
    student.Group = updatedData.Group
    student.Email = updatedData.Email
    
    if err := config.GetDB().Save(&student).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al actualizar estudiante")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Estudiante actualizado exitosamente", student)
}

// DeleteStudent maneja DELETE /api/students/:student_id
func DeleteStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("student_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var student models.Student
    if err := config.GetDB().First(&student, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    if err := config.GetDB().Delete(&student).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al eliminar estudiante")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Estudiante eliminado exitosamente", nil)
}