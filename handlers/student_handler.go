package handlers

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"
)

// CreateStudent godoc
// @Summary      Crear un nuevo estudiante
// @Description  Registra un nuevo estudiante en el sistema
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student  body      models.Student  true  "Información del estudiante"
// @Success      201      {object}  utils.SuccessResponse{data=models.Student}
// @Failure      400      {object}  utils.ErrorResponse
// @Failure      500      {object}  utils.ErrorResponse
// @Router       /students [post]
func CreateStudent(c *gin.Context) {
    var student models.Student
    
    if err := c.ShouldBindJSON(&student); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    if err := config.GetDB().Create(&student).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear el estudiante")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusCreated, "Estudiante creado exitosamente", student)
}

// GetAllStudents godoc
// @Summary      Listar todos los estudiantes
// @Description  Obtiene la lista completa de estudiantes registrados
// @Tags         students
// @Produce      json
// @Success      200  {array}   models.Student
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /students [get]
func GetAllStudents(c *gin.Context) {
    var students []models.Student
    
    if err := config.GetDB().Find(&students).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener estudiantes")
        return
    }
    
    c.JSON(http.StatusOK, students)
}

// GetStudent godoc
// @Summary      Obtener un estudiante por ID
// @Description  Obtiene la información detallada de un estudiante específico
// @Tags         students
// @Produce      json
// @Param        student_id  path      int  true  "ID del estudiante"
// @Success      200         {object}  models.Student
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Router       /students/{student_id} [get]
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

// UpdateStudent godoc
// @Summary      Actualizar un estudiante
// @Description  Actualiza la información de un estudiante existente
// @Tags         students
// @Accept       json
// @Produce      json
// @Param        student_id  path      int             true  "ID del estudiante"
// @Param        student     body      models.Student  true  "Información actualizada del estudiante"
// @Success      200         {object}  utils.SuccessResponse{data=models.Student}
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Failure      500         {object}  utils.ErrorResponse
// @Router       /students/{student_id} [put]
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
    
    student.Name = updatedData.Name
    student.Group = updatedData.Group
    student.Email = updatedData.Email
    
    if err := config.GetDB().Save(&student).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al actualizar estudiante")
        return
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Estudiante actualizado exitosamente", student)
}

// DeleteStudent godoc
// @Summary      Eliminar un estudiante
// @Description  Elimina un estudiante del sistema (también elimina sus calificaciones por CASCADE)
// @Tags         students
// @Produce      json
// @Param        student_id  path      int  true  "ID del estudiante"
// @Success      200         {object}  utils.SuccessResponse
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Failure      500         {object}  utils.ErrorResponse
// @Router       /students/{student_id} [delete]
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