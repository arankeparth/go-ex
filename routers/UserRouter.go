package routers

import (
	"go-ex/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpUserRouters(r *gin.Engine) {
	// Routes for CRUD operations
	r.POST("/users", handlers.CreateUsersHandler)

	r.GET("/users", handlers.GetUsersHandler)

	r.PUT("/users/:id", handlers.UpdateUsersHandler)

	r.DELETE("/users/:id", handlers.DeleteUsersHandler)
}
