package main

import (
	"context"
	"net/http"

	"go-ex/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Initialize MongoDB
	ctx := context.Background()
	db.InitializeDatabase(ctx)

	// Routes for CRUD operations
	r.POST("/users", func(c *gin.Context) {
		var user bson.M
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := db.CreateDocument(ctx, "exampleDB", "users", user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"insertedID": result.InsertedID})
	})

	r.GET("/users", func(c *gin.Context) {
		filter := bson.M{} // Fetch all documents
		users, err := db.ReadDocuments(ctx, "exampleDB", "users", filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var update bson.M
		if err := c.ShouldBindJSON(&update); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		filter := bson.M{"_id": id} // Replace with ObjectID conversion if required
		result, err := db.UpdateDocument(ctx, "exampleDB", "users", filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"matchedCount": result.MatchedCount, "modifiedCount": result.ModifiedCount})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		filter := bson.M{"_id": id} // Replace with ObjectID conversion if required

		result, err := db.DeleteDocument(ctx, "exampleDB", "users", filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deletedCount": result.DeletedCount})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
