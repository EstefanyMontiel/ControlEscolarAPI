package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"
)

// CreateGrade maneja POST /api/grades
func CreateGrade(c *gin.Context) {
    var grade models.Grade
    
    if err := c.ShouldBindJSON(&grade); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    // Verificar que el estudiante existe
    var student models.Student
    if err := config.GetDB().First(&student, grade.StudentID).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    // Verificar que la materia existe
    var subject models.Subject
    if err := config.GetDB().First(&subject, grade.SubjectID).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Materia no encontrada")
        return
    }
    
    if err := config.GetDB().Create(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear la calificación")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusCreated, "Calificación creada exitosamente", grade)
}

// UpdateGrade maneja PUT /api/grades/:grade_id
func UpdateGrade(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("grade_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var grade models.Grade
    if err := config.GetDB().First(&grade, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Calificación no encontrada")
        return
    }
    
    var updatedData models.Grade
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    grade.Grade = updatedData.Grade
    
    if err := config.GetDB().Save(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al actualizar calificación")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Calificación actualizada exitosamente", grade)
}

// DeleteGrade maneja DELETE /api/grades/:grade_id
func DeleteGrade(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("grade_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var grade models.Grade
    if err := config.GetDB().First(&grade, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Calificación no encontrada")
        return
    }
    
    if err := config.GetDB().Delete(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al eliminar calificación")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Calificación eliminada exitosamente", nil)
}

// GetGradeByStudentAndSubject maneja GET /api/grades/:grade_id/student/:student_id
func GetGradeByStudentAndSubject(c *gin.Context) {
    gradeID, err := strconv.Atoi(c.Param("grade_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID de calificación inválido")
        return
    }
    
    studentID, err := strconv.Atoi(c.Param("student_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID de estudiante inválido")
        return
    }
    
    var grade models.Grade
    if err := config.GetDB().
        Preload("Student").
        Preload("Subject").
        Where("grade_id = ? AND student_id = ?", gradeID, studentID).
        First(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Calificación no encontrada")
        return
    }
    
    c.JSON(http.StatusOK, grade)
}

// GetStudentGrades maneja GET /api/grades/student/:student_id
func GetStudentGrades(c *gin.Context) {
    studentID, err := strconv.Atoi(c.Param("student_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID de estudiante inválido")
        return
    }
    
    // Verificar que el estudiante existe
    var student models.Student
    if err := config.GetDB().First(&student, studentID).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    var grades []models.Grade
    if err := config.GetDB().
        Preload("Subject").
        Where("student_id = ?", studentID).
        Find(&grades).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener calificaciones")
        return
    }
    
    c.JSON(http.StatusOK, grades)
}
