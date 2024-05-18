package database

import (
	"context"
	"log"

	"github.com/vanoraco/vshop/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/* var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIDIsNotValid   = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("cannot add product to cart")
	ErrCantRemoveItem     = errors.New("cannot remove item from cart")
	ErrCantGetItem        = errors.New("cannot get item from cart ")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
) */

func AddProductToShop(ctx context.Context, prodCollection, shopCollection *mongo.Collection, productID primitive.ObjectID, ownerID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return ErrCantFindProduct
	}
	var productList []models.Product
	err = searchfromdb.All(ctx, &productList)
	if err != nil {
		log.Println(err)
		return ErrCantDecodeProducts
	}

	id, err := primitive.ObjectIDFromHex(ownerID)
	if err != nil {
		log.Println(err)
		return ErrUserIDIsNotValid
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "shop_products", Value: bson.D{{Key: "$each", Value: productList}}}}}}
	_, err = shopCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return ErrCantUpdateUser
	}
	return nil
}
