package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vanoraco/vshop/database"
	"github.com/vanoraco/vshop/models"
	generate "github.com/vanoraco/vshop/tokens"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var OwnerCollection *mongo.Collection = database.UserData(database.Client, "Owners")
var ProductCollection *mongo.Collection = database.ProductData(database.Client, "Products")
var ShopCollection *mongo.Collection = database.ShopData(database.Client, "Shops")
var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userpassword string, givenpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(givenpassword), []byte(userpassword))
	valid := true
	msg := ""
	if err != nil {
		msg = "Login Or Password is Incorrect"
		valid = false
	}
	return valid, msg
}

func SignUpUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		}
		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()
		token, refreshtoken, _ := generate.TokenGenerator(*user.Email, *user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshtoken
		user.UserCart = make([]models.ProductUser, 0)
		user.Address_Details = make([]models.Address, 0)
		user.Order_Status = make([]models.Order, 0)
		_, inserterr := UserCollection.InsertOne(ctx, user)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusCreated, "Successfully Signed Up!!")
	}
}

func SignUpOwners() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var owner models.Owner

		// Bind form data to the owner struct
		if err := c.ShouldBind(&owner); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the user already exists by email
		count, err := OwnerCollection.CountDocuments(ctx, bson.M{"owner_email": owner.Owner_Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		// Check if the user already exists by phone
		count, err = OwnerCollection.CountDocuments(ctx, bson.M{"owner_phone": owner.Owner_Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Phone is already in use"})
			return
		}

		// Hash the password
		hashedPassword := HashPassword(*owner.Owner_Password)
		owner.Owner_Password = &hashedPassword

		// Set additional fields
		owner.Created_At = time.Now()
		owner.Updated_At = time.Now()
		owner.ID = primitive.NewObjectID()
		owner.Owner_ID = owner.ID.Hex()

		// Generate tokens
		token, refreshtoken, _ := generate.TokenGenerator(*owner.Owner_Email, *owner.Owner_FirstName, *owner.Owner_LastName, owner.Owner_ID)
		owner.Token = &token
		owner.Refresh_Token = &refreshtoken

		// Handle file upload to Cloudinary
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
			return
		}
		defer file.Close()

		// Create Cloudinary instance
		cld, err := cloudinary.NewFromParams("djwh3jl5j", "547786729594556", "hEZvcmeQ3nFN6IKtPsqpHoFlSK0")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
			return
		}

		// Upload file to Cloudinary
		uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "shop_images"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to Cloudinary"})
			return
		}

		// Set the image URL to the owner
		owner.Owner_Img = uploadResult.SecureURL

		// Insert owner into the database
		_, inserterr := OwnerCollection.InsertOne(ctx, owner)
		if inserterr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not created"})
			return
		}

		c.JSON(http.StatusCreated, "Successfully Signed Up!!")
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var user models.User
		var founduser models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "login or password incorrect"})
			return
		}
		PasswordIsValid, msg := VerifyPassword(*user.Password, *founduser.Password)
		defer cancel()
		if !PasswordIsValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			fmt.Println(msg)
			return
		}
		token, refreshToken, _ := generate.TokenGenerator(*founduser.Email, *founduser.First_Name, *founduser.Last_Name, founduser.User_ID)
		defer cancel()
		generate.UpdateAllTokens(token, refreshToken, founduser.User_ID)
		c.JSON(http.StatusOK, founduser)

	}
}

func ProductViewerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var products models.Product
		defer cancel()
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		products.Product_ID = primitive.NewObjectID()
		_, anyerr := ProductCollection.InsertOne(ctx, products)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, "Successfully added our Product Admin!!")
	}
}

func AddShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		ownerID := c.Query("owner_id")
		if ownerID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing owner_id parameter"})
			return
		}

		var owner models.Owner
		err := OwnerCollection.FindOne(ctx, bson.M{"owner_id": ownerID}).Decode(&owner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find owner"})
			return
		}

		var shop models.Shop
		if err := c.ShouldBind(&shop); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		count, err := ShopCollection.CountDocuments(ctx, bson.M{"shop_id": ownerID})

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You Have Opened Shop already"})
			return
		}

		// Bind JSON request body to Shop struct

		shop.Shop_ID = &ownerID

		shop.Shop_Owner = owner.Owner_FirstName
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
			return
		}
		defer file.Close()

		// Create Cloudinary instance
		cld, err := cloudinary.NewFromParams("djwh3jl5j", "547786729594556", "hEZvcmeQ3nFN6IKtPsqpHoFlSK0")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
			return
		}

		// Upload file to Cloudinary
		uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "shop_images"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to Cloudinary"})
			return
		}

		// Set the image URL to the product
		shop.Image = uploadResult.SecureURL

		// Generate a new ObjectID for the shop

		// Insert the shop into the database
		_, err = ShopCollection.InsertOne(ctx, shop)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add shop"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, "Shop added successfully")
	}
	/* return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var shop models.Shop
		defer cancel()
		if err := c.BindJSON(&shop); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ownerQueryId := c.Query("ownerID")
		shop.Shop_ID = &ownerQueryId

		ownerID, err := primitive.ObjectIDFromHex(ownerQueryId)
		if err != nil {
			log.Fatal(err)
		}
		filter := bson.M{"_id": ownerID}

		// Define options to include only the "Owner_FirstName" field
		options := options.FindOne().SetProjection(bson.M{"Owner_FirstName": 1})

		// Find the owner document based on ID and apply projection
		var result bson.M
		if err := OwnerCollection.FindOne(context.Background(), filter, options).Decode(&result); err != nil {
			log.Fatal(err)
		}

		shop.Shop_Owner = result["owner_firstname"].(string)
		_, anyerr := ShopCollection.InsertOne(ctx, shop)
		if anyerr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Not Created"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, "Shop Successfully Added")
	} */
}

/* func AddProductsToShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get shop ID from query parameter
		shopID := c.Query("owner_id")
		if shopID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing shop_id parameter"})
			return
		}

		// Fetch shop details from the shop collection
		var shop models.Shop
		err := ShopCollection.FindOne(ctx, bson.M{"shop_id": shopID}).Decode(&shop)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find shop"})
			return
		}

		// Bind JSON request body to []ProductUser
		var products models.Product
		if err := c.BindJSON(&products); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageURL, err := uploadImageToImgBB(c, "image")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to ImgBB"})
			return
		}

		// Set the image URL to the product
		products.Image = imageURL
		// Append the new products to the existing shop products
		shop.Shop_Products = append(shop.Shop_Products, products)

		// Update the shop document in the MongoDB collection
		update := bson.M{"$set": bson.M{"shop_products": shop.Shop_Products}}
		_, err = ShopCollection.UpdateOne(ctx, bson.M{"shop_id": shopID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shop"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, "Products added to shop successfully")
	}
} */

/* func AddProductsToShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get shop ID from query parameter
		shopID := c.Query("owner_id")
		if shopID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing shop_id parameter"})
			return
		}

		// Fetch shop details from the shop collection
		var shop models.Shop
		err := ShopCollection.FindOne(ctx, bson.M{"shop_id": shopID}).Decode(&shop)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find shop"})
			return
		}

		// Bind form file to Product struct
		var product models.Product
		if err := c.Bind(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Upload image to ImgBB and get the URL
		imageURL, err := uploadImageToImgBB(c, product.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		// Set the image URL to the product
		product.Image = imageURL

		// Append the new product to the existing shop products
		shop.Shop_Products = append(shop.Shop_Products, product)

		// Update the shop document in the MongoDB collection
		update := bson.M{"$set": bson.M{"shop_products": shop.Shop_Products}}
		_, err = ShopCollection.UpdateOne(ctx, bson.M{"shop_id": shopID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shop"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, "Products added to shop successfully")
	}
} */

/*
	 func AddProductsToShop() gin.HandlerFunc {
		return func(c *gin.Context) {
			// Create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
			defer cancel()

			// Get shop ID from query parameter
			shopID := c.Query("owner_id")
			if shopID == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Missing shop_id parameter"})
				return
			}

			// Fetch shop details from the shop collection
			var shop models.Shop
			err := ShopCollection.FindOne(ctx, bson.M{"shop_id": shopID}).Decode(&shop)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find shop"})
				return
			}

			// Bind form file to Product struct
			var product models.Product
			if err := c.Bind(&product); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Get the file from the form data
			file, _, err := c.Request.FormFile("file")
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
				return
			}
			defer file.Close()

			// Create Cloudinary instance
			cld, err := cloudinary.NewFromParams("djwh3jl5j", "547786729594556", "hEZvcmeQ3nFN6IKtPsqpHoFlSK0")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
				return
			}

			// Upload file to Cloudinary
			uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "product_images"})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to Cloudinary"})
				return
			}

			// Set the image URL to the product
			product.Image = uploadResult.SecureURL

			// Append the new product to the existing shop products
			shop.Shop_Products = append(shop.Shop_Products, product)

			// Update the shop document in the MongoDB collection
			update := bson.M{"$set": bson.M{"shop_products": shop.Shop_Products}}
			_, err = ShopCollection.UpdateOne(ctx, bson.M{"shop_id": shopID}, update)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shop"})
				return
			}

			// Respond with success message
			c.JSON(http.StatusOK, "Products added to shop successfully")
		}
	}
*/
func AddProductsToShop() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Get shop ID from query parameter
		shopID := c.Query("owner_id")
		if shopID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing shop_id parameter"})
			return
		}

		// Fetch shop details from the shop collection
		var shop models.Shop
		err := ShopCollection.FindOne(ctx, bson.M{"shop_id": shopID}).Decode(&shop)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find shop"})
			return
		}

		// Bind form data to Product struct
		var product models.Product
		if err := c.ShouldBind(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a new ObjectID for the product
		product.Product_ID = primitive.NewObjectID()

		// Get the file from the form data
		file, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file"})
			return
		}
		defer file.Close()

		// Create Cloudinary instance
		cld, err := cloudinary.NewFromParams("djwh3jl5j", "547786729594556", "hEZvcmeQ3nFN6IKtPsqpHoFlSK0")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize Cloudinary"})
			return
		}

		// Upload file to Cloudinary
		uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "product_images"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image to Cloudinary"})
			return
		}

		// Set the image URL to the product
		product.Image = uploadResult.SecureURL

		// Update the shop document in the MongoDB collection using $push to append the new product
		update := bson.M{"$push": bson.M{"shop_products": product}}
		_, err = ShopCollection.UpdateOne(ctx, bson.M{"shop_id": shopID}, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update shop"})
			return
		}

		// Respond with success message
		c.JSON(http.StatusOK, gin.H{"message": "Product added to shop successfully"})
	}
}

/* func uploadImageToImgBB(c *gin.Context, formField string) (string, error) {

	type ImgBBResponse struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	// Parse form
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		return "", err
	}

	// Get the file from the form data
	file, _, err := c.Request.FormFile(formField)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file field
	part, err := writer.CreateFormFile("image", "image.png")
	if err != nil {
		return "", err
	}

	// Copy file content to part
	if _, err := io.Copy(part, file); err != nil {
		return "", err
	}

	// Close multipart writer
	if err := writer.Close(); err != nil {
		return "", err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.imgbb.com/1/upload", body)
	if err != nil {
		return "", err
	}

	// Set request headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth("ec63f76d0d1c91469a942f9c44cad0f1", "")

	// Perform the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Decode ImgBB API response
	var imgBBRes ImgBBResponse
	if err := json.NewDecoder(res.Body).Decode(&imgBBRes); err != nil {
		return "", err
	}

	return imgBBRes.Data.URL, nil
} */

func SearchProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var productlist []models.Product
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		cursor, err := ProductCollection.Find(ctx, bson.D{{}})
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "Someting Went Wrong Please Try After Some Time")
			return
		}
		err = cursor.All(ctx, &productlist)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		defer cursor.Close(ctx)
		if err := cursor.Err(); err != nil {
			// Don't forget to log errors. I log them really simple here just
			// to get the point across.
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer cancel()
		c.IndentedJSON(200, productlist)

	}
}

func SearchProductByQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		var searchproducts []models.Product
		queryParam := c.Query("name")
		if queryParam == "" {
			log.Println("query is empty")
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"Error": "Invalid Search Index"})
			c.Abort()
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		searchquerydb, err := ProductCollection.Find(ctx, bson.M{"product_name": bson.M{"$regex": queryParam}})
		if err != nil {
			c.IndentedJSON(404, "something went wrong in fetching the dbquery")
			return
		}
		err = searchquerydb.All(ctx, &searchproducts)
		if err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid")
			return
		}
		defer searchquerydb.Close(ctx)
		if err := searchquerydb.Err(); err != nil {
			log.Println(err)
			c.IndentedJSON(400, "invalid request")
			return
		}
		defer cancel()
		c.IndentedJSON(200, searchproducts)
	}
}
