package routes

import (
	"Tamer/model"
	"Tamer/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func DeletePost(c *gin.Context) {
	isAdmin, err := c.Request.Cookie("isAdmin")
	if err != nil || isAdmin.Value != "true" {
		c.JSON(http.StatusForbidden, "Only admin is able to interact with posts!")
		return
	}

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

	_, err = services.DeleteArticle(idParam, collection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
