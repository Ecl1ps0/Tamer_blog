package config

import (
	"Tamer/model"
	"github.com/gin-gonic/gin"
)

func SetCollectionContext(db *model.DBCollection) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("collection", db)
		c.Next()
	}
}
