package models

import (
    "gorm.io/gorm"
)

// Grade representa una calificación en el sistema
type Grade struct {
    GradeID   int      `gorm:"primaryKey;autoIncrement" json:"grade_id" example:"1"`
    StudentID int      `gorm:"not null;index" json:"student_id" binding:"required,min=1" example:"1"`
    SubjectID int      `gorm:"not null;index" json:"subject_id" binding:"required,min=1" example:"1"`
    Grade     float64  `gorm:"type:decimal(5,2);not null" json:"grade" binding:"required,min=0,max=100" example:"95.5"`
}

func (Grade) TableName() string {
    return "grades"
}

func MigrateGrade(db *gorm.DB) error {
    return db.AutoMigrate(&Grade{})
}

func AddForeignKeys(db *gorm.DB) error {
    if err := db.Exec(`
        ALTER TABLE grades 
        ADD CONSTRAINT fk_grades_student 
        FOREIGN KEY (student_id) 
        REFERENCES students(student_id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
    `).Error; err != nil {
        if !db.Migrator().HasConstraint(&Grade{}, "fk_grades_student") {
            return err
        }
    }
    
    if err := db.Exec(`
        ALTER TABLE grades 
        ADD CONSTRAINT fk_grades_subject 
        FOREIGN KEY (subject_id) 
        REFERENCES subjects(subject_id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
    `).Error; err != nil {
        if !db.Migrator().HasConstraint(&Grade{}, "fk_grades_subject") {
            return err
        }
    }
    
    return nil
}