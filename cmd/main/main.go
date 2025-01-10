package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/paytm-temp/cms-backend/pkg/api/v1/handlers"
    "github.com/paytm-temp/cms-backend/pkg/api/v1/middleware"
)

func main() {
    r := gin.Default()

    // CORS middleware
    r.Use(func(c *gin.Context) {
        origin := c.Request.Header.Get("Origin")
        if origin == "http://localhost:3000" || origin == "http://localhost:3001" {
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
            log.Printf("[DEBUG] Setting CORS headers for origin: %s", origin)
        }
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, X-User-Role")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Max-Age", "3600")
        c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

        if c.Request.Method == "OPTIONS" {
            log.Printf("[DEBUG] Handling OPTIONS request from origin: %s", origin)
            c.AbortWithStatus(204)
            return
        }

        log.Printf("[DEBUG] Incoming request: %s %s from origin: %s", c.Request.Method, c.Request.URL.Path, origin)
        c.Next()
    })

    // Add role middleware
    r.Use(middleware.RoleMiddleware())

    // Register routes
    handlers.RegisterRoutes(r)

    // Start server
    log.Printf("[INFO] Starting server on :8080")
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
