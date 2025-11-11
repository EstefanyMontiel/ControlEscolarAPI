package main

import (
    "log"
    
    "github.com/gin-gonic/gin"
	"ControlEscolar/config"
	"ControlEscolar/models"
    "ControlEscolar/routes"
)

func main() {
    // Inicializar base de datos
    config.InitDatabase()
    
    // Ejecutar migraciones
    db := config.GetDB()
    if err := models.MigrateStudent(db); err != nil {
        log.Fatal("Error en migración de estudiantes:", err)
    }
    if err := models.MigrateSubject(db); err != nil {
        log.Fatal("Error en migración de materias:", err)
    }
    if err := models.MigrateGrade(db); err != nil {
        log.Fatal("Error en migración de calificaciones:", err)
    }
    
    log.Println("Migraciones ejecutadas exitosamente")
    
    // Configurar Gin
    router := gin.Default()
    
    // Middleware para CORS (opcional)
    router.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })
    
    // Configurar rutas
    routes.SetupRoutes(router)
    
    // Ruta de prueba
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "API de Control Escolar - Sistema funcionando correctamente",
            "version": "1.0.0",
        })
    })
    
    // Iniciar servidor
    log.Println("Servidor iniciado en http://localhost:8080")
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Error al iniciar servidor:", err)
    }
}