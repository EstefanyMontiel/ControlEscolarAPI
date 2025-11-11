package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
	"ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"

)

// CreateSubject maneja POST /api/subjects
func CreateSubject(c *gin.Context) {
    var subject models.Subject
    
    if err := c.ShouldBindJSON(&subject); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    if err := config.GetDB().Create(&subject).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear la materia")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusCreated, "Materia creada exitosamente", subject)
}

// GetSubject maneja GET /api/subjects/:subject_id
func GetSubject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("subject_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var subject models.Subject
    if err := config.GetDB().First(&subject, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Materia no encontrada")
        return
    }
    
    c.JSON(http.StatusOK, subject)
}

// UpdateSubject maneja PUT /api/subjects/:subject_id
func UpdateSubject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("subject_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var subject models.Subject
    if err := config.GetDB().First(&subject, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Materia no encontrada")
        return
    }
    
    var updatedData models.Subject
    if err := c.ShouldBindJSON(&updatedData); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    subject.Name = updatedData.Name
    
    if err := config.GetDB().Save(&subject).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al actualizar materia")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Materia actualizada exitosamente", subject)
}

// DeleteSubject maneja DELETE /api/subjects/:subject_id
func DeleteSubject(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("subject_id"))
    if err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "ID inválido")
        return
    }
    
    var subject models.Subject
    if err := config.GetDB().First(&subject, id).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Materia no encontrada")
        return
    }
    
    if err := config.GetDB().Delete(&subject).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al eliminar materia")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Materia eliminada exitosamente", nil)
}
