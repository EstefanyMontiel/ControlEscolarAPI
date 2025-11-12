package models

import (
    "gorm.io/gorm"
)

// Subject representa una materia en el sistema
type Subject struct {
    SubjectID int    `gorm:"primaryKey;autoIncrement" json:"subject_id"`
    Name      string `gorm:"type:varchar(100);unique;not null" json:"name" binding:"required,min=2,max=100"`
}

// TableName especifica el nombre de la tabla
func (Subject) TableName() string {
    return "subjects"
}

// MigrateSubject ejecuta las migraciones para la tabla de materias
func MigrateSubject(db *gorm.DB) error {
    return db.AutoMigrate(&Subject{})
}