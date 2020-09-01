package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"herosi/controllers"
	"herosi/models"
)

func main() {
	router := gin.Default()

	db := models.SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.Use(cors.Default())
	router.Use(static.Serve("/", static.LocalFile("./AngularApp", true)))

	api := router.Group("/api")
	{
		api.GET("/heroes", controllers.GetHeroes)
		api.GET("/heroes/:id", controllers.GetHero)
		api.POST("/heroes", controllers.AddHero)
		api.PUT("/heroes/:id", controllers.PutHero)
		api.DELETE("/heroes/:id", controllers.DeleteHero)
	}

	router.Run()
}
