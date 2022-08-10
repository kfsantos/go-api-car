package controllers

import (
	"go-api-car/database"
	"go-api-car/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CarList(c *gin.Context) {
	var items []models.Car
	database.DB.Order("id asc").Find(&items)
	c.JSON(http.StatusOK, items)
}

func CarCreate(c *gin.Context) {
	var item models.Car

	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty fields, no data to save!"})
		return
	}

	if err := database.DB.Create(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error saving data!"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func CarGetById(c *gin.Context) {
	id := c.Params.ByName("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "A valid identifier is required!"})
		return
	}
	var item models.Car
	err = database.DB.Find(&item, newId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Item nor found!"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func CarUpdate(c *gin.Context) {
	id := c.Params.ByName("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "A valid identifier is required!"})
		return
	}

	var item models.Car
	err = database.DB.First(&item, newId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no item in DB!"})
		return
	}

	err = c.ShouldBindJSON(&item)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error when performing binding!"})
		return
	}
	database.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

func CarDelete(c *gin.Context) {
	id := c.Params.ByName("id")
	newId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "A valid identifier is required!"})
		return
	}

	var item models.Car
	err = database.DB.First(&item, newId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There is no item in DB!"})
		return
	}
	database.DB.Delete(&item, newId)
	c.JSON(http.StatusOK, gin.H{
		"data": "Record deleted successfully!"})

}
