package handlers

import (
    "log"
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/utils"
)

// CreateGrade godoc
// @Summary      Crear una nueva calificación
// @Description  Registra una nueva calificación para un estudiante en una materia
// @Tags         grades
// @Accept       json
// @Produce      json
// @Param        grade  body      models.CreateGradeRequest  true  "Información de la calificación"
// @Success      201    {object}  utils.SuccessResponse{data=models.GradeResponse}
// @Failure      400    {object}  utils.ErrorResponse
// @Failure      404    {object}  utils.ErrorResponse
// @Failure      500    {object}  utils.ErrorResponse
// @Router       /grades [post]
func CreateGrade(c *gin.Context) {
    var request models.CreateGradeRequest
    
    if err := c.ShouldBindJSON(&request); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    // Verificar que el estudiante existe
    var student models.Student
    if err := config.GetDB().First(&student, request.StudentID).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Estudiante no encontrado")
        return
    }
    
    // Verificar que la materia existe
    var subject models.Subject
    if err := config.GetDB().First(&subject, request.SubjectID).Error; err != nil {
        utils.RespondWithError(c, http.StatusNotFound, "Materia no encontrada")
        return
    }
    
    // Crear el registro de calificación
    grade := models.Grade{
        StudentID: request.StudentID,
        SubjectID: request.SubjectID,
        Grade:     request.Grade,
    }
    
    if err := config.GetDB().Create(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al crear la calificación")
        return
    }
    
    // Preparar respuesta con información completa
    response := models.GradeResponse{
        GradeID:   grade.GradeID,
        StudentID: grade.StudentID,
        SubjectID: grade.SubjectID,
        Grade:     grade.Grade,
        Student: &models.StudentBasic{
            StudentID: student.StudentID,
            Name:      student.Name,
            Group:     student.Group,
            Email:     student.Email,
        },
        Subject: &models.SubjectBasic{
            SubjectID: subject.SubjectID,
            Name:      subject.Name,
        },
    }
    
    utils.RespondWithSuccess(c, http.StatusCreated, "Calificación creada exitosamente", response)
}

// UpdateGrade godoc
// @Summary      Actualizar una calificación
// @Description  Actualiza el valor de una calificación existente
// @Tags         grades
// @Accept       json
// @Produce      json
// @Param        grade_id  path      int                        true  "ID de la calificación"
// @Param        grade     body      models.UpdateGradeRequest  true  "Nueva calificación"
// @Success      200       {object}  utils.SuccessResponse{data=models.GradeResponse}
// @Failure      400       {object}  utils.ErrorResponse
// @Failure      404       {object}  utils.ErrorResponse
// @Failure      500       {object}  utils.ErrorResponse
// @Router       /grades/{grade_id} [put]
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
    
    var request models.UpdateGradeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, "Datos inválidos: "+err.Error())
        return
    }
    
    // Actualizar solo el campo grade
    grade.Grade = request.Grade
    
    if err := config.GetDB().Save(&grade).Error; err != nil {
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al actualizar calificación")
        return
    }
    
    // Obtener información completa para la respuesta
    var student models.Student
    var subject models.Subject
    config.GetDB().First(&student, grade.StudentID)
    config.GetDB().First(&subject, grade.SubjectID)
    
    response := models.GradeResponse{
        GradeID:   grade.GradeID,
        StudentID: grade.StudentID,
        SubjectID: grade.SubjectID,
        Grade:     grade.Grade,
        Student: &models.StudentBasic{
            StudentID: student.StudentID,
            Name:      student.Name,
            Group:     student.Group,
            Email:     student.Email,
        },
        Subject: &models.SubjectBasic{
            SubjectID: subject.SubjectID,
            Name:      subject.Name,
        },
    }
    
    utils.RespondWithSuccess(c, http.StatusOK, "Calificación actualizada exitosamente", response)
}

// DeleteGrade godoc
// @Summary      Eliminar una calificación
// @Description  Elimina una calificación del sistema
// @Tags         grades
// @Produce      json
// @Param        grade_id  path      int  true  "ID de la calificación"
// @Success      200       {object}  utils.SuccessResponse
// @Failure      400       {object}  utils.ErrorResponse
// @Failure      404       {object}  utils.ErrorResponse
// @Failure      500       {object}  utils.ErrorResponse
// @Router       /grades/{grade_id} [delete]
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

// GetGradeByStudentAndSubject godoc
// @Summary      Obtener calificación específica
// @Description  Obtiene una calificación específica de un estudiante por grade_id y student_id
// @Tags         grades
// @Produce      json
// @Param        grade_id    path      int  true  "ID de la calificación"
// @Param        student_id  path      int  true  "ID del estudiante"
// @Success      200         {object}  models.GradeResponse
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Router       /grades/{grade_id}/student/{student_id} [get]
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
    
    // Buscar la calificación
    var grade models.Grade
    if err := config.GetDB().
        Where("grade_id = ? AND student_id = ?", gradeID, studentID).
        First(&grade).Error; err != nil {
        log.Printf("Error buscando calificación: %v", err)
        utils.RespondWithError(c, http.StatusNotFound, "Calificación no encontrada")
        return
    }
    
    // Obtener información del estudiante y materia
    var student models.Student
    var subject models.Subject
    config.GetDB().First(&student, grade.StudentID)
    config.GetDB().First(&subject, grade.SubjectID)
    
    response := models.GradeResponse{
        GradeID:   grade.GradeID,
        StudentID: grade.StudentID,
        SubjectID: grade.SubjectID,
        Grade:     grade.Grade,
        Student: &models.StudentBasic{
            StudentID: student.StudentID,
            Name:      student.Name,
            Group:     student.Group,
            Email:     student.Email,
        },
        Subject: &models.SubjectBasic{
            SubjectID: subject.SubjectID,
            Name:      subject.Name,
        },
    }
    
    c.JSON(http.StatusOK, response)
}

// GetStudentGrades godoc
// @Summary      Obtener todas las calificaciones de un estudiante
// @Description  Obtiene todas las calificaciones registradas para un estudiante específico
// @Tags         grades
// @Produce      json
// @Param        student_id  path      int  true  "ID del estudiante"
// @Success      200         {array}   models.GradeResponse
// @Failure      400         {object}  utils.ErrorResponse
// @Failure      404         {object}  utils.ErrorResponse
// @Failure      500         {object}  utils.ErrorResponse
// @Router       /grades/student/{student_id} [get]
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
    
    // Obtener todas las calificaciones del estudiante
    var grades []models.Grade
    if err := config.GetDB().
        Where("student_id = ?", studentID).
        Find(&grades).Error; err != nil {
        log.Printf("Error obteniendo calificaciones: %v", err)
        utils.RespondWithError(c, http.StatusInternalServerError, "Error al obtener calificaciones")
        return
    }
    
    // Preparar respuestas con información completa
    var responses []models.GradeResponse
    for _, grade := range grades {
        var subject models.Subject
        config.GetDB().First(&subject, grade.SubjectID)
        
        response := models.GradeResponse{
            GradeID:   grade.GradeID,
            StudentID: grade.StudentID,
            SubjectID: grade.SubjectID,
            Grade:     grade.Grade,
            Student: &models.StudentBasic{
                StudentID: student.StudentID,
                Name:      student.Name,
                Group:     student.Group,
                Email:     student.Email,
            },
            Subject: &models.SubjectBasic{
                SubjectID: subject.SubjectID,
                Name:      subject.Name,
            },
        }
        responses = append(responses, response)
    }
    
    c.JSON(http.StatusOK, responses)
}