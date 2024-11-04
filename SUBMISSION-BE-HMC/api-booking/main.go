package main

import (
	"api-booking/database"
	"api-booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	db := database.ConnectToDb()
	routes.Routes(r, db)

	r.Run(":8000")

}