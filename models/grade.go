package models

import (
    "gorm.io/gorm"
)

// Grade representa una calificación en el sistema
type Grade struct {
    GradeID   int     `gorm:"primaryKey;autoIncrement" json:"grade_id"`
    StudentID int     `gorm:"not null;index" json:"student_id" binding:"required,min=1"`
    SubjectID int     `gorm:"not null;index" json:"subject_id" binding:"required,min=1"`
    Grade     float64 `gorm:"type:decimal(5,2);not null" json:"grade" binding:"required,min=0,max=100"`
    
    // Relaciones - SOLO para consultas, no afectan la migración
    Student   *Student `gorm:"-" json:"student,omitempty"`
    Subject   *Subject `gorm:"-" json:"subject,omitempty"`
}

// TableName especifica el nombre de la tabla
func (Grade) TableName() string {
    return "grades"
}

// MigrateGrade ejecuta las migraciones para la tabla de calificaciones
func MigrateGrade(db *gorm.DB) error {
    return db.AutoMigrate(&Grade{})
}

// AddForeignKeys agrega las llaves foráneas DESPUÉS de crear todas las tablas
func AddForeignKeys(db *gorm.DB) error {
    // Agregar llave foránea para student_id
    if err := db.Exec(`
        ALTER TABLE grades 
        ADD CONSTRAINT fk_grades_student 
        FOREIGN KEY (student_id) 
        REFERENCES students(student_id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
    `).Error; err != nil {
        // Si ya existe, ignorar el error
        if !db.Migrator().HasConstraint(&Grade{}, "fk_grades_student") {
            return err
        }
    }
    
    // Agregar llave foránea para subject_id
    if err := db.Exec(`
        ALTER TABLE grades 
        ADD CONSTRAINT fk_grades_subject 
        FOREIGN KEY (subject_id) 
        REFERENCES subjects(subject_id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
    `).Error; err != nil {
        // Si ya existe, ignorar el error
        if !db.Migrator().HasConstraint(&Grade{}, "fk_grades_subject") {
            return err
        }
    }
    
    return nil
}