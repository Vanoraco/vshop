package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/vanoraco/vshop/controllers"
	"github.com/vanoraco/vshop/database"
	"github.com/vanoraco/vshop/middleware"
	"github.com/vanoraco/vshop/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	shopApp := controllers.ShopNewApplication(database.ProductData(database.Client, "Products"), database.OwnerData(database.Client, "Owners"))

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"}, // Replace with your frontend origin
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.POST("/addtoshop", shopApp.AddToShop())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))
}
