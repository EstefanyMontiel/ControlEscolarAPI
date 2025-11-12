package models

// CreateGradeRequest representa la petición para crear una calificación
type CreateGradeRequest struct {
    StudentID int     `json:"student_id" binding:"required,min=1" example:"1"`
    SubjectID int     `json:"subject_id" binding:"required,min=1" example:"1"`
    Grade     float64 `json:"grade" binding:"required,min=0,max=100" example:"95.5"`
}

// UpdateGradeRequest representa la petición para actualizar una calificación
type UpdateGradeRequest struct {
    Grade float64 `json:"grade" binding:"required,min=0,max=100" example:"98.0"`
}

// GradeResponse representa la respuesta de una calificación con información completa
type GradeResponse struct {
    GradeID   int             `json:"grade_id" example:"1"`
    StudentID int             `json:"student_id" example:"1"`
    SubjectID int             `json:"subject_id" example:"1"`
    Grade     float64         `json:"grade" example:"95.5"`
    Student   *StudentBasic   `json:"student,omitempty"`
    Subject   *SubjectBasic   `json:"subject,omitempty"`
}

// StudentBasic información básica de estudiante
type StudentBasic struct {
    StudentID int    `json:"student_id" example:"1"`
    Name      string `json:"name" example:"María García"`
    Group     string `json:"group" example:"5A"`
    Email     string `json:"email" example:"maria@escuela.com"`
}

// SubjectBasic información básica de materia
type SubjectBasic struct {
    SubjectID int    `json:"subject_id" example:"1"`
    Name      string `json:"name" example:"Matemáticas"`
}