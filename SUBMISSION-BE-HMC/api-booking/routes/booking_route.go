package routes

import (
	"api-booking/handler"
	"api-booking/helpers"
	"api-booking/repositories"
	"api-booking/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookingRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repoBooking := repositories.NewBookingRepository(db)
	repoRoom := repositories.NewRoomRepository(db)
	svc := services.NewBookingService(repoBooking, repoRoom)
	handler := handler.NewBookingHandler(svc)

	auth := r.Use(helpers.JWTMiddleware())
	auth.GET("", helpers.AuthorizeMiddleware("admin"), handler.GetAllBooking)
	auth.GET("/:id", helpers.AuthorizeMiddleware("admin"), handler.GetBookingById)
	auth.POST("", handler.Create)
	auth.PUT("/:id", helpers.AuthorizeMiddleware("admin"), handler.Update)
	auth.DELETE("/:id", helpers.AuthorizeMiddleware("admin"), handler.Delete)
}