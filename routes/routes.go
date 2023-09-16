package routes

import (
	"github.com/Farhatkiani/ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.AddProduct())
	incomingRoutes.GET("/users/productview", controllers.ProductView())
	incomingRoutes.GET("/user/search", controllers.Search())

}
