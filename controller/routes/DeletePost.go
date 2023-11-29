package routes

import (
	"Tamer/model"
	"Tamer/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DeletePost(c *gin.Context) {
	logger := log.New(c.Writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	isAdmin, err := c.Request.Cookie("isAdmin")
	if err != nil || isAdmin.Value != "true" {
		c.JSON(http.StatusForbidden, "Only admin is able to interact with posts!")
		logger.Fatalf("Rights error: %v\n", err)
	}

	collectionInterface, exists := c.Get("collection")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get collection from the context!"})
		logger.Fatal("The context of collection is empty!")
	}

	collection, ok := collectionInterface.(*model.DBCollection)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse collection after getting the context!"})
		logger.Fatal("Failed to assert the type to *model.DBCollection!")
	}

	idParam := c.Param("id")

	_, err = services.DeleteArticle(idParam, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the post"})
		logger.Fatalf("Failed to delete the post: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
