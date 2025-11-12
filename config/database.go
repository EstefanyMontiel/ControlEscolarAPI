package config

import (
    "fmt"
    "log"
    "os"
    
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDatabase inicializa la conexión a MySQL
func InitDatabase() {
    var err error
    
    // Configuración de conexión MySQL
    // Formato: usuario:contraseña@tcp(host:puerto)/nombre_base_datos?charset=utf8mb4&parseTime=True&loc=Local
    
    // Obtener credenciales desde variables de entorno (recomendado)
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "")
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbName := getEnv("DB_NAME", "control_escolar")
    
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)
    
    // Conectar a MySQL
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    
    if err != nil {
        log.Fatal("Error al conectar a la base de datos MySQL:", err)
    }
    
    log.Println("✅ Conexión a MySQL establecida exitosamente")
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
    return DB
}

// getEnv obtiene variable de entorno o usa valor por defecto
func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}