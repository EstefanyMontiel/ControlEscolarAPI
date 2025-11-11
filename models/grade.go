package models

import (
    "gorm.io/gorm"
)

// Grade representa una calificación en el sistema
type Grade struct {
    GradeID   int     `gorm:"primaryKey;autoIncrement" json:"grade_id"`
    StudentID int     `gorm:"not null" json:"student_id" binding:"required"`
    SubjectID int     `gorm:"not null" json:"subject_id" binding:"required"`
    Grade     float64 `gorm:"type:decimal(5,2);not null" json:"grade" binding:"required,min=0,max=100"`
    
    // Relaciones
    Student   Student `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE" json:"student,omitempty"`
    Subject   Subject `gorm:"foreignKey:SubjectID;constraint:OnDelete:CASCADE" json:"subject,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Grade) TableName() string {
    return "grades"
}

// Migrate ejecuta las migraciones para la tabla de calificaciones
func MigrateGrade(db *gorm.DB) error {
    return db.AutoMigrate(&Grade{})
}