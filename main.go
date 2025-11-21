package main

import (
    "log"
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/routes"
    
    _ "ControlEscolar/docs"
)

// @title           API de Control Escolar
// @version         1.0
// @description     API REST para la gestión de estudiantes, materias y calificaciones en un sistema escolar

// @contact.name   Estefany Montiel
// @contact.email  estefany.montiel@example.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8082
// @BasePath  /api

// @schemes http
func main() {
    // Cargar variables de entorno
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️  Advertencia: No se encontró archivo .env")
    } else {
        log.Println("✅ Archivo .env cargado correctamente")
    }
    
    // Inicializar base de datos
    config.InitDatabase()
    
    db := config.GetDB()
    
    // Ejecutar migraciones
    log.Println("📦 Ejecutando migraciones...")
    if err := models.MigrateStudent(db); err != nil {
        log.Fatal("❌ Error en migración de estudiantes:", err)
    }
    if err := models.MigrateSubject(db); err != nil {
        log.Fatal("❌ Error en migración de materias:", err)
    }
    if err := models.MigrateGrade(db); err != nil {
        log.Fatal("❌ Error en migración de calificaciones:", err)
    }
    
    // Agregar llaves foráneas
    if err := models.AddForeignKeys(db); err != nil {
        log.Println("⚠️  Advertencia al agregar llaves foráneas:", err)
    }
    
    log.Println("✅ Base de datos lista")
    
    // Configurar Gin
    router := gin.Default()
    router.SetTrustedProxies(nil)
    
    // Middleware para CORS
    router.Use(CORSMiddleware())
    
    // Configurar rutas de la API
    routes.SetupRoutes(router)
    
    // Ruta principal
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "API de Control Escolar",
            "version": "1.0.0",
            "status":  "running",
            "docs":    "http://localhost:8082/swagger/index.html",
        })
    })
    
    // Ruta para Swagger UI
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    
    // Obtener puerto
    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }
    
    // Iniciar servidor
    log.Printf("🚀 Servidor iniciado en http://localhost:%s\n", port)
    log.Printf("📚 Documentación Swagger: http://localhost:%s/swagger/index.html\n", port)
    log.Printf("🔍 Prueba las rutas: http://localhost:%s/api/students\n", port)
    
    if err := router.Run(":" + port); err != nil {
        log.Fatal("❌ Error al iniciar servidor:", err)
    }
}

// CORSMiddleware configura CORS
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}