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
	cam := "/heroes"

	{
		api.GET(cam, controllers.GetHeroes)
		api.GET(cam+"/:id", controllers.GetHero)
		api.POST(cam, controllers.AddHero)
		api.PUT(cam+"/:id", controllers.PutHero)
		api.DELETE(cam+"/:id", controllers.DeleteHero)
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./AngularApp/index.html")
	})

	router.Run()
}
