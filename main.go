package main

import (
	controllers "sport-teams/controllers"
	"sport-teams/models"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	api := r.Group("/api")
	{
		teams := api.Group("/teams")
		{
			teams.GET("", controllers.FindTeams)
			teams.POST("", controllers.CreateTeam)
			teams.GET("/:id", controllers.FindTeam)
			teams.PATCH("/:id", controllers.UpdateTeam)
			teams.DELETE("/:id", controllers.DeleteTeam)
		}
	}

	r.GET("/metrics", prometheusHandler())

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
