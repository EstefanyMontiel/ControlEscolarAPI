package utils

import (
    "github.com/gin-gonic/gin"
)

// ErrorResponse estructura para respuestas de error
type ErrorResponse struct {
    Error   string `json:"error"`
    Message string `json:"message,omitempty"`
}

// SuccessResponse estructura para respuestas exitosas
type SuccessResponse struct {
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

// RespondWithError envía una respuesta de error
func RespondWithError(c *gin.Context, code int, message string) {
    c.JSON(code, ErrorResponse{
        Error:   "Error",
        Message: message,
    })
}

// RespondWithSuccess envía una respuesta exitosa
func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
    c.JSON(code, SuccessResponse{
        Message: message,
        Data:    data,
    })
}