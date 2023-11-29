package routes

import (
	"Tamer/model"
	"Tamer/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetAllPosts(c *gin.Context) {
	logger := log.New(c.Writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	collectionInterface, exists := c.Get("collection")
	if !exists {
		logger.Fatal("The context of collection is empty!")
	}

	collection, ok := collectionInterface.(*model.DBCollection)
	if !ok {
		logger.Fatal("Failed to assert the type to *model.DBCollection!")
	}

	posts, err := services.GetAllPosts(collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.Fatalf("Getting posts error: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
