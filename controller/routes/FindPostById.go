package routes

import (
	"Tamer/model"
	"Tamer/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindPostBYId(c *gin.Context) {
	logger := log.New(c.Writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	collectionInterface, exists := c.Get("collection")
	if !exists {
		logger.Fatal("The context of collection is empty!")
	}

	collection, ok := collectionInterface.(*model.DBCollection)
	if !ok {
		logger.Fatal("Failed to assert the type to *model.DBCollection!")
	}

	idParam := c.Param("id")

	post, err := services.FindBYId(idParam, collection)
	if err != nil {
		logger.Fatalf("Failed to find the post by id: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}
