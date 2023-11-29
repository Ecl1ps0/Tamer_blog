package controller

import (
	"Tamer/config"
	"Tamer/controller/routes"
	"Tamer/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer(db *model.DBCollection, errorLogger *log.Logger) {
	router := gin.New()

	gin.SetMode(gin.ReleaseMode)

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		errorLogger.Fatalf("Set trusted proxies error: %v\n", err)
	}

	router.Use(config.SetCollectionContext(db))

	router.GET("/", routes.GetAllPosts)
	router.DELETE("/delete/:id", routes.DeletePost)
	router.GET("/get/:id", routes.FindPostBYId)
	router.POST("/create", routes.CreatePost)
	router.PUT("/update/:id", routes.UpdatePost)
	router.POST("/authorize", routes.CheckIsAdmin)

	if err := router.Run(":8080"); err != nil {
		errorLogger.Fatalf("Start server error: %v\n", err)
	}
}
