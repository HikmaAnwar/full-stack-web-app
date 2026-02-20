package routes

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes the API routes
func SetupRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
	{
		// Add your routes here
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
