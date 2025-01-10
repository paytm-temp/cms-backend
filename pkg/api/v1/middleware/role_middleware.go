package middleware

import (
    "github.com/gin-gonic/gin"
)

const (
    RoleHeader = "X-User-Role"
)

type Role string

const (
    RoleRequestor   Role = "REQUESTOR"
    RoleCaseManager Role = "CASE_MANAGER"
    RoleResolver    Role = "RESOLVER"
)

func RoleMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        role := Role(c.GetHeader(RoleHeader))
        
        // Allow all roles to access the API
        if role == "" {
            role = RoleRequestor // Default role
        }
        
        c.Set("userRole", role)
        c.Next()
    }
}
