package routes

import (
	"Tamer/config"
	"Tamer/model"
	"Tamer/services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func UpdatePost(c *gin.Context) {
	logger := log.New(c.Writer, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	isAdmin, err := c.Request.Cookie("isAdmin")
	if err != nil || isAdmin.Value != "true" {
		c.JSON(http.StatusForbidden, "Only admin is able to interact with posts!")
		logger.Fatalf("Rights error: %v\n", err)
	}

	collectionInterface, exists := c.Get("collection")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse collection after getting the context!"})
		logger.Fatal("The context of collection is empty!")
	}

	collection, ok := collectionInterface.(*model.DBCollection)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get collection from the context!"})
		logger.Fatal("Failed to assert the type to *model.DBCollection!")
	}

	idParam := c.Param("id")

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse the form!"})
		logger.Fatalf("Failed to parse the form: %v\n", err)
	}

	title := c.Request.PostFormValue("title")
	textContent := c.Request.PostFormValue("textContent")

	image, _, err := c.Request.FormFile("imageContent")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fail to load image!"})
		logger.Fatal("Fail to load image!")
	}
	defer image.Close()

	imageBytes, err := config.ConvertImgToBytes(image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to convert image to bytes!"})
		logger.Fatalf("Image converting error: %v\n", err)
	}

	base64Img := base64.StdEncoding.EncodeToString(imageBytes)

	updatedArticle := &model.Article{
		Title:        title,
		TextContent:  textContent,
		ImageContent: base64Img,
		CreateAt:     uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		UpdatedAt:    uint64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	_, updateErr := services.UpdateArticle(idParam, updatedArticle, collection)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the post in database!"})
		logger.Fatalf("Update error: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post successfully updated!"})
}
