package middleware

import (
	usersConcertsModel "comedians/src/core/usersConcerts/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredPermissions []string) func(c *gin.Context) {
	return func(c *gin.Context) {
		contextRoles, _ := c.Get("userRoles")
		roles := contextRoles.([]usersConcertsModel.Role)

		for _, requiredPermission := range requiredPermissions {
			isContains := false

			for _, role := range roles {
				for _, permission := range role.Permissions {
					if permission.Title == requiredPermission {
						isContains = true
					}
				}
			}

			if !isContains {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
		c.Next()
	}
}
