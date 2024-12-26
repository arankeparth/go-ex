package handlers

import (
	"go-ex/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUsersHandler(c *gin.Context) {
	var user bson.M
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.CreateDocument(c, "exampleDB", "users", user)
	log.Printf("Creating new user: %+v", user)
	if err != nil {
		log.Printf("Error while creating user: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("User Created successfully, ID: %s", result.InsertedID)
	c.JSON(http.StatusOK, gin.H{"insertedID": result.InsertedID})
}

func GetUsersHandler(c *gin.Context) {
	// for i := 0; i < 1000; i++ {
	// 	_ = make([]byte, 1024*1024) // Allocating memory
	// }
	filter := bson.M{} // Fetch all documents
	users, err := db.ReadDocuments(c, "exampleDB", "users", filter)
	if err != nil {
		log.Printf("Error while fetching all users: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("All users fetched successfully.")
	c.JSON(http.StatusOK, users)
}

func UpdateUsersHandler(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var update bson.M
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Updating User with id: %s, Updating following properties: %+v", id.String(), update)
	result, err := db.UpdateDocument(c, "exampleDB", "users", id, update)
	if err != nil {
		log.Printf("Error while updating User with id: %s, %s", id.String(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.MatchedCount > 0 {
		log.Printf("User with id: %s updated successfully.", id.String())
	} else {
		log.Printf("User not found with id: %s", id.String())
	}

	c.JSON(http.StatusOK, gin.H{"matchedCount": result.MatchedCount, "modifiedCount": result.ModifiedCount})
}

func DeleteUsersHandler(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": id} // Replace with ObjectID conversion if required

	result, err := db.DeleteDocument(c, "exampleDB", "users", filter)
	if err != nil {
		log.Printf("Error while deleting User with id: %s, %s", id.String(), err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount > 0 {
		log.Printf("User with id: %s deleted successfully.", id.String())
	} else {
		log.Printf("User not found with id: %s", id.String())
	}

	c.JSON(http.StatusOK, gin.H{"deletedCount": result.DeletedCount})
}
