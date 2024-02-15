package main

import (
	apicontrollers "sisbus/controllers/apiControllers"

	"sisbus/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	//r.Use(cors.Default())

	// Middleware Logger
	r.Use(logger.SetLogger())
	r.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"*"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET,POST,PUT,DELETE"},
	}))
	//Studi
	r.GET("/api/Studi", apicontrollers.IndexStudi)
	r.GET("/api/Studi/:id", apicontrollers.ShowStudi)
	r.POST("/api/Studi", apicontrollers.CreateStudi)
	r.PUT("/api/Studi/:id", apicontrollers.UpdateStudi)
	r.DELETE("/api/Studi/dellet/:id", apicontrollers.DeleteStudi)
	//public
	r.Run()
}
