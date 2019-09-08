package controllers

import (
	"App/database"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

// Controller is a generic struct
type Controller struct {
}

// Collection returns the collection from the database
func (Controller) Collection(collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

// Log is a logger
func (Controller) Log(c *gin.Context) *log.Entry {
	contextualizedLogger := log.WithFields(
		log.Fields{
			"route": c.Request.RequestURI,
		})
	return contextualizedLogger
}
