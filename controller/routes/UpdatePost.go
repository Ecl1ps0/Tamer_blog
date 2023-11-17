package routes

import (
	"Tamer/config"
	"Tamer/model"
	"Tamer/services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UpdatePost(c *gin.Context) {
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

	postID, parseErr := strconv.ParseUint(idParam, 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse the form!"})
		return
	}

	title := c.Request.PostFormValue("title")
	textContent := c.Request.PostFormValue("textContent")

	image, _, err := c.Request.FormFile("imageContent")
	if err != nil {
		log.Print("Fail to load image!")
		return
	}
	defer image.Close()

	imageBytes, err := config.ConvertImgToBytes(image)
	if err != nil {
		log.Printf("Image converting error: %v\n", err)
	}

	base64Img := base64.StdEncoding.EncodeToString(imageBytes)

	updatedArticle := &model.Article{
		Title:        title,
		TextContent:  textContent,
		ImageContent: base64Img,
		CreateAt:     uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		UpdatedAt:    uint64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	_, updateErr := services.UpdateArticle(postID, updatedArticle, collection)
	if updateErr != nil {
		log.Printf("Update error: %v\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post successfully updated!"})
}
