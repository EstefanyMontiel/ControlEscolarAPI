package config

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase inicializa la conexión a la base de datos
func InitDatabase() {
    var err error
    
    // Usar SQLite para simplicidad (puedes cambiar a PostgreSQL o MySQL)
    DB, err = gorm.Open(sqlite.Open("school.db"), &gorm.Config{})
    
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }
    
    log.Println("Conexión a la base de datos establecida exitosamente")
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
    return DB
}