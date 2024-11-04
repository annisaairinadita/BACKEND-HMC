package routes

import (
	"api-booking/handler"
	"api-booking/helpers"
	"api-booking/repositories"
	"api-booking/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoomRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewRoomRepository(db)
	svc := services.NewRoomService(repo)
	handler := handler.NewRoomHandler(svc)

	auth := r.Use(helpers.JWTMiddleware())
	auth.GET("", handler.GetAllRoom)
	auth.GET("/:id", handler.GetRoomById)
	auth.POST("", helpers.AuthorizeMiddleware("admin"), handler.Create)
	auth.GET("/number/:number", handler.GetRoomByRoomNumber)
	auth.PUT("/:id", helpers.AuthorizeMiddleware("admin"), handler.Update)
	auth.DELETE("/:id", helpers.AuthorizeMiddleware("admin"), handler.Delete)
}