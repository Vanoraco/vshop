package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name"`
	Last_Name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *int               `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id,omitempty"`
	Product_Name *string            `json:"product_name" form:"product_name"`
	Price        *uint64            `json:"price" form:"price"`
	Rating       *uint8             `json:"rating,omitempty" form:"rating,omitempty"`
	Image        string             `json:"image,omitempty"`
}

type Shop struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Shop_Name     *string            `json:"shop_name" form:"shop_name"`
	Shop_Owner    *string            `json:"shop_owner" form:"shop_owner"`
	Shop_Category *string            `json:"shop_category" form:"shop_category"`
	Image         string             `json:"image"`
	Rating        *uint8             `json:"rating" form:"rating"`
	Shop_ID       *string            `json:"shop_id"`
	ShopProducts  []Product          `json:"shop_products,omitempty" bson:"shop_products,omitempty"`
}
type Owner struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	Owner_FirstName *string            `json:"owner_firstname" bson:"owner_firstname" form:"owner_firstname"`
	Owner_LastName  *string            `json:"owner_lastname" bson:"owner_lastname" form:"owner_lastname"`
	Owner_Img       string             `json:"owner_img" bson:"owner_img" form:"owner_img"`
	Owner_Email     *string            `json:"owner_email" bson:"owner_email" form:"owner_email"`
	Owner_Password  *string            `json:"owner_password" bson:"owner_password" form:"owner_password"`
	Owner_Phone     *string            `json:"owner_phone" bson:"owner_phone" form:"owner_phone"`
	Token           *string            `json:"token" bson:"token"`
	Refresh_Token   *string            `json:"refresh_token" bson:"refresh_token"`
	Created_At      time.Time          `json:"created_at" bson:"created_at"`
	Updated_At      time.Time          `json:"updated_at" bson:"updated_at" `
	Owner_ID        string             `json:"owner_id" bson:"owner_id" `
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        int                `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}

type Address struct {
	Address_ID primitive.ObjectID `json:"_id" bson:"_id"`
	House      *string            `json:"house" bson:"name"`
	Street     *string            `json:"street" bson:"street"`
	City       *string            `json:"city" bson:"city"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price          int                `json:"price" bson:"price"`
	Discount       *int               `json:"discount" bson:"discount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
