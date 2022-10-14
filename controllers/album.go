package controllers

import (
	"net/http"
	"strconv"
	"web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func FindAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func FindAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	c.JSON(http.StatusOK, gin.H{"data": album})
}

func CreateAlbum(c *gin.Context) {
	var input models.CreateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var price float64
	if f, err := strconv.ParseFloat(input.Price, 64); err == nil {
		price = f
	}

	album := models.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  price,
	}
	models.DB.Create(&album)
	c.JSON(http.StatusOK, gin.H{"data": album})
}

func UpdateAlbum(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}

	var input models.UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&album).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": album})
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album
	if err := models.DB.Where("id = ?", c.Param("id")).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&album)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func ApplyAlbumRouter(router *gin.Engine) *gin.Engine {
	router.GET("/albums", FindAlbums)
	router.POST("/albums", CreateAlbum)
	router.GET("/albums/:id", FindAlbum)
	router.PATCH("/albums/:id", UpdateAlbum)
	router.DELETE("/albums/:id", DeleteAlbum)
	return router
}
