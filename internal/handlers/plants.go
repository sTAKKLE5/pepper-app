package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pepper-app/internal/database"
	"pepper-app/internal/models"
	"strconv"
	"strings"
	"time"
)

func CreatePlant(c *gin.Context) {
	var plantRequestBody struct {
		Species      string
		Cultivar     string
		PlantingDate string
		IsCross      string
	}

	if err := c.Bind(&plantRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body", "err": err.Error()})
		return
	}

	isCrossBool, err := strconv.ParseBool(strings.ToLower(plantRequestBody.IsCross))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse is_cross"})
		return
	}

	plantingDate, _ := time.Parse("2006-01-02", plantRequestBody.PlantingDate)
	plant := models.Plant{
		Species:      strings.ToLower(plantRequestBody.Species),
		Cultivar:     plantRequestBody.Cultivar,
		PlantingDate: plantingDate,
		IsCross:      isCrossBool,
	}

	_, err = database.DB.NamedExec(`INSERT INTO plants (species, cultivar, planting_date, is_cross) 
		VALUES (:species, :cultivar, :planting_date, :is_cross)`,
		map[string]interface{}{
			"species":       plant.Species,
			"cultivar":      plant.Cultivar,
			"planting_date": plant.PlantingDate,
			"is_cross":      plant.IsCross,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create plant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plant created successfully"})
}

func GetPlants(c *gin.Context) {
	var plants []models.Plant

	err := database.DB.Select(&plants, "SELECT * FROM plants")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plants"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plants": plants})
}

func GetPlantByID(c *gin.Context) {
	id := c.Param("id")
	var plant models.Plant

	err := database.DB.Get(&plant, "SELECT * FROM plants WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get plant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"plant": plant})
}

func UpdatePlantByID(c *gin.Context) {
	id := c.Param("id")
	var plantRequestBody struct {
		Species      *string
		Cultivar     *string
		PlantingDate *string
		IsCross      *bool
	}
	if err := c.Bind(&plantRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body", "err": err.Error()})
		return
	}

	// Prepare the update query
	updateQuery := `UPDATE plants SET species = :species, cultivar = :cultivar, planting_date = :planting_date, is_cross = :is_cross WHERE id = :id`

	// Prepare the parameters
	params := map[string]interface{}{
		"id":            id,
		"species":       plantRequestBody.Species,
		"cultivar":      plantRequestBody.Cultivar,
		"planting_date": plantRequestBody.PlantingDate,
		"is_cross":      plantRequestBody.IsCross,
	}

	// Execute the update query
	_, err := database.DB.NamedExec(updateQuery, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update plant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plant updated successfully"})
}

func DeletePlantByID(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM plants WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete plant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Plant deleted successfully"})
}
