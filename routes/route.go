package routes

import (
	"dashboard-ecommerce-team2/infra"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRoutes(ctx infra.ServiceContext) *gin.Engine {
	
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := r.Group("/products")
	{
		router.POST("/", ctx.Ctl.Product.CreateProductController)
		router.GET("/", ctx.Ctl.Product.GetAllProductsController)
		router.GET("/:id", ctx.Ctl.Product.GetProductByIDController)
		router.DELETE("/:id", ctx.Ctl.Product.DeleteProductController)
	}

	return r
}
