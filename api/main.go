package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	var newAlbums []album
	var isAlbumFound bool = false

	for _, album := range albums {
		if album.ID == id {
			isAlbumFound = true
		}

		if album.ID != id {
			newAlbums = append(newAlbums, album)
		}
	}
	albums = newAlbums

	if isAlbumFound {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}
}

func putAlbum(c *gin.Context) {
	id := c.Param("id")
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	newAlbum.ID = id

	for i, album := range albums {
		if album.ID == id {
			albums[i] = newAlbum
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "album updated"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("albums/:id", deleteAlbum)
	router.PUT("albums/:id", putAlbum)

	router.Run("localhost:8080")
}
