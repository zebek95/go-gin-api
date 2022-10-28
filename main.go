package main

import (
	"github.com/gin-gonic/gin"
)

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	router := gin.Default()

	albumsRouterGroup := router.Group("/albums")
	{
		albumsRouterGroup.GET("/", getAlbums)
		albumsRouterGroup.GET("/:id", getAlbum)
		albumsRouterGroup.POST("/", postAlbums)
	}

	router.Run("localhost:8080")
}
