package routes

import (
	"Tamer/config"
	"Tamer/model"
	"Tamer/services"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {
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

	article := model.Article{
		ID:           primitive.NewObjectID(),
		Title:        title,
		TextContent:  textContent,
		ImageContent: base64Img,
		CreateAt:     uint64(time.Now().UnixNano() / int64(time.Millisecond)),
		UpdatedAt:    uint64(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	_, err = services.CreateArticle(article, collection)
	if err != nil {
		log.Printf("Database saving error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the article to the database!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article created successfully!"})
}
