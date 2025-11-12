package main

import (
    "log"
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "ControlEscolar/config"
    "ControlEscolar/models"
    "ControlEscolar/routes"
)

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
        
    // Paso 1: Crear tablas padre (sin llaves foráneas)
    log.Println("📦 Creando tablas padre...")
    if err := models.MigrateStudent(db); err != nil {
        log.Fatal("❌ Error en migración de estudiantes:", err)
    }
    log.Println("✅ Tabla students creada")
    
    if err := models.MigrateSubject(db); err != nil {
        log.Fatal("❌ Error en migración de materias:", err)
    }
    log.Println("✅ Tabla subjects creada")
    
    // Paso 2: Crear tabla hija (sin llaves foráneas todavía)
    log.Println("📦 Creando tabla grades...")
    if err := models.MigrateGrade(db); err != nil {
        log.Fatal("❌ Error en migración de calificaciones:", err)
    }
    log.Println("✅ Tabla grades creada")
    
    // Paso 3: Agregar llaves foráneas
    log.Println("🔗 Agregando llaves foráneas...")
    if err := models.AddForeignKeys(db); err != nil {
        log.Println("⚠️  Advertencia al agregar llaves foráneas:", err)
    } else {
        log.Println("✅ Llaves foráneas agregadas")
    }
    
    log.Println("✅ Todas las migraciones ejecutadas exitosamente")
    
    // Configurar Gin
    router := gin.Default()
    
    // Configurar proxies confiables (quita el warning)
    router.SetTrustedProxies(nil)
    
    // Middleware para CORS
    router.Use(CORSMiddleware())
    
    // Configurar rutas
    routes.SetupRoutes(router)
    
    // Ruta de prueba
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "API de Control Escolar - Sistema funcionando correctamente",
            "version": "1.0.0",
            "database": "MySQL",
        })
    })
    
    // Obtener puerto desde variable de entorno
    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }
    
    // Iniciar servidor
    log.Printf("🚀 Servidor iniciado en http://localhost:%s\n", port)
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