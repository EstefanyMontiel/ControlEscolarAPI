package models

import (
    "gorm.io/gorm"
)

// Subject representa una materia en el sistema
type Subject struct {
    SubjectID int    `gorm:"primaryKey;autoIncrement" json:"subject_id" example:"1"`
    Name      string `gorm:"type:varchar(100);unique;not null" json:"name" binding:"required,min=2,max=100" example:"Matemáticas"`
}

func (Subject) TableName() string {
    return "subjects"
}

func MigrateSubject(db *gorm.DB) error {
    return db.AutoMigrate(&Subject{})
}