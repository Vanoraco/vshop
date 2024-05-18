package controllers

import (
	"context"
	"net/http"
	"time"

	//"github.com/vanoraco/vshop/database"
	"github.com/vanoraco/vshop/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type ShopApplication struct {
	prodCollection *mongo.Collection
	shopCollection *mongo.Collection
}

func ShopNewApplication(prodCollection, shopCollection *mongo.Collection) *ShopApplication {
	return &ShopApplication{
		prodCollection: prodCollection,
		shopCollection: shopCollection,
	}
}

func (shopApp *ShopApplication) AddToShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product

		// Bind the JSON request body to the product struct
		if err := c.BindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a new ObjectID for the product
		product.Product_ID = primitive.NewObjectID()

		// Get the owner ID from the query parameter
		ownerQueryID := c.Query("ownerID")
		if ownerQueryID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "owner id is empty"})
			return
		}

		// Create a context with a timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Convert owner ID string to ObjectID
		ownerObjectID, err := primitive.ObjectIDFromHex(ownerQueryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid owner id"})
			return
		}

		// Define the filter to find the owner by its ObjectID
		filter := bson.D{{Key: "_id", Value: ownerObjectID}}

		// Define the update operation to push the product details into the owner_product array field
		update := bson.D{
			{Key: "$push", Value: bson.D{
				{Key: "shop_products", Value: product},
			}},
		}

		// Execute the update operation
		_, err = shopApp.shopCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, gin.H{"message": "Successfully added to the shop"})
	}
}
