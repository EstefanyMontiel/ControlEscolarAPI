package models

import (
    "gorm.io/gorm"
)

// Student representa un estudiante en el sistema
type Student struct {
    StudentID int    `gorm:"primaryKey;autoIncrement" json:"student_id" example:"1"`
    Name      string `gorm:"type:varchar(100);not null" json:"name" binding:"required,min=2,max=100" example:"María García"`
    Group     string `gorm:"type:varchar(10);not null" json:"group" binding:"required,min=1,max=10" example:"5A"`
    Email     string `gorm:"type:varchar(100);unique;not null" json:"email" binding:"required,email" example:"maria.garcia@escuela.com"`
}

func (Student) TableName() string {
    return "students"
}

func MigrateStudent(db *gorm.DB) error {
    return db.AutoMigrate(&Student{})
}