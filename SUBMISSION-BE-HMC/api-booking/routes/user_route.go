package routes

import (
	"api-booking/handler"
	"api-booking/helpers"
	"api-booking/repositories"
	"api-booking/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	svc := services.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)

	auth := group.Use(helpers.JWTMiddleware())
	auth.GET("/:id", handler.GetProfile)
	auth.GET("", helpers.AuthorizeMiddleware("admin"), handler.GetAllUser)
	auth.DELETE("/:id", handler.Delete)
	auth.GET("/me", handler.Me)
	auth.PUT("/:id", handler.Update)
}