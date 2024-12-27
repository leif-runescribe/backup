// routes/user_routes.go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leif-runescribe/westeros-roster/models"
	"gorm.io/gorm"
)

func SetupUserRoutes(r *gin.Engine, db *gorm.DB) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", func(c *gin.Context) {
			var user models.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := db.Create(&user)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			c.JSON(http.StatusCreated, user)
		})

		userRoutes.GET("/", func(c *gin.Context) {
			var users []models.User
			result := db.Find(&users)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}
			c.JSON(http.StatusOK, users)
		})

		userRoutes.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var user models.User
			result := db.First(&user, id)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, user)
		})

		userRoutes.PUT("/:id", func(c *gin.Context) {
			id := c.Param("id")
			var user models.User
			result := db.First(&user, id)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		})

		userRoutes.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")
			result := db.Delete(&models.User{}, id)
			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
		})
	}
}
