package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"
)

// CreateSubject godoc
// @Summary      Crear una nueva materia
// @Description  Registra una nueva materia en el sistema
// @Tags         subjects
// @Accept       json
// @Produce      json
// @Param        subject  body      models.Subject  true  "Información de la materia"
// @Success      201      {object}  utils.SuccessResponse{data=models.Subject}
// @Failure      400      {object}  utils.ErrorResponse
// @Failure      500      {object}  utils.ErrorResponse
// @Router       /subjects [post]
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

// GetSubject godoc
// @Summary      Obtener una materia por ID
// @Description  Obtiene la información de una materia específica
// @Tags         subjects
// @Produce      json
// @Param        subject_id  path      int  true  "ID de la materia"
// @Success      200         {object}  models.Subject
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Router       /subjects/{subject_id} [get]
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

// UpdateSubject godoc
// @Summary      Actualizar una materia
// @Description  Actualiza la información de una materia existente
// @Tags         subjects
// @Accept       json
// @Produce      json
// @Param        subject_id  path      int             true  "ID de la materia"
// @Param        subject     body      models.Subject  true  "Información actualizada de la materia"
// @Success      200         {object}  utils.SuccessResponse{data=models.Subject}
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Failure      500         {object}  utils.ErrorResponse
// @Router       /subjects/{subject_id} [put]
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

// DeleteSubject godoc
// @Summary      Eliminar una materia
// @Description  Elimina una materia del sistema (también elimina sus calificaciones por CASCADE)
// @Tags         subjects
// @Produce      json
// @Param        subject_id  path      int  true  "ID de la materia"
// @Success      200         {object}  utils.SuccessResponse
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Failure      500         {object}  utils.ErrorResponse
// @Router       /subjects/{subject_id} [delete]
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