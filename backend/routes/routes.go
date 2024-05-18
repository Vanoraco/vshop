package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vanoraco/vshop/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUpUsers())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/owner/signupowner", controllers.SignUpOwners())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearchProduct())
	incomingRoutes.GET("/users/search", controllers.SearchProductByQuery())
	incomingRoutes.POST("/owners/addshop", controllers.AddShop())
	incomingRoutes.POST("/owners/addproducttoshop", controllers.AddProductsToShop())
}
