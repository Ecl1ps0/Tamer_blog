package routes

import (
	"Tamer/model"
	"Tamer/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindPostBYId(c *gin.Context) {
	collectionInterface, exists := c.Get("collection")
	if !exists {
		log.Print("The context of collection is empty!")
		return
	}

	collection, ok := collectionInterface.(*model.DBCollection)
	if !ok {
		log.Print("Failed to assert the type to *model.DBCollection!")
		return
	}

	idParam := c.Param("id")

	post, err := services.FindBYId(idParam, collection)
	if err != nil {
		log.Printf("Failed to find the post by id: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}