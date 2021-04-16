package controllers

import (
	"net/http"
	models "sport-teams/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /teams
// Get all teams
func FindTeams(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var teams []models.Team
	db.Find(&teams)

	c.JSON(http.StatusOK, gin.H{"data": teams})
}

// POST /teams
// Create new teams
func CreateTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreateTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Team
	team := models.Team{Name: input.Name, Coach: input.Coach, Email: input.Email}
	db.Create(&team)

	c.JSON(http.StatusOK, gin.H{"data": team})

}

// GET /teams/:id
// Find a team
func FindTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var team models.Team
	if err := db.Where("id = ?", c.Param("id")).First(&team).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

// PATCH /teams/:id
// Update a team
func UpdateTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var team models.Team
	if err := db.Where("id = ?", c.Param("id")).First(&team).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateTeamInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&team).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": team})
}

// DELETE /teams/:id
// Delete a team
func DeleteTeam(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var team models.Team
	if err := db.Where("id = ?", c.Param("id")).First(&team).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&team)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
