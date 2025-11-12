package models

import (
    "gorm.io/gorm"
)

// Student representa un estudiante en el sistema
type Student struct {
    StudentID int    `gorm:"primaryKey;autoIncrement" json:"student_id"`
    Name      string `gorm:"type:varchar(100);not null" json:"name" binding:"required,min=2,max=100"`
    Group     string `gorm:"type:varchar(10);not null" json:"group" binding:"required,min=1,max=10"`
    Email     string `gorm:"type:varchar(100);unique;not null" json:"email" binding:"required,email"`
}

// TableName especifica el nombre de la tabla
func (Student) TableName() string {
    return "students"
}

// MigrateStudent ejecuta las migraciones para la tabla de estudiantes
func MigrateStudent(db *gorm.DB) error {
    return db.AutoMigrate(&Student{})
}