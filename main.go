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

	r.GET("/metrics", prometheusHandler())

	r.GET("/teams", controllers.FindTeams)

	r.POST("/teams", controllers.CreateTeam) // create

	r.GET("/teams/:id", controllers.FindTeam) // find by id

	r.PATCH("/teams/:id", controllers.UpdateTeam) // update by id

	r.DELETE("/teams/:id", controllers.DeleteTeam) // delete by id

	r.Run()
}
