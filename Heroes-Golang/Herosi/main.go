package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	heroc "herosi/controllers"
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
	cam := "/heroitems"

	{
		api.GET(cam, heroc.GetHeroes)
		api.GET(cam+"/:id", heroc.GetHero)
		api.POST(cam, heroc.AddHero)
		api.PUT(cam+"/:id", heroc.PutHero)
		api.DELETE(cam+"/:id", heroc.DeleteHero)
	}

	router.NoRoute(func(c *gin.Context) {
		c.File("./AngularApp/index.html")
	})

	router.Run()
}
