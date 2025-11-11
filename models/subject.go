package models

import (
    "gorm.io/gorm"
)

// Subject representa una materia en el sistema
type Subject struct {
    SubjectID int    `gorm:"primaryKey;autoIncrement" json:"subject_id"`
    Name      string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
}

// TableName especifica el nombre de la tabla
func (Subject) TableName() string {
    return "subjects"
}

// Migrate ejecuta las migraciones para la tabla de materias
func MigrateSubject(db *gorm.DB) error {
    return db.AutoMigrate(&Subject{})
}