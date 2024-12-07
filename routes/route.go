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

	authRoutes(r, ctx)

	productRoutes := r.Group("/products")
	{
		productRoutes.POST("/", ctx.Ctl.Product.CreateProductController)
	}

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.GET("/", ctx.Ctl.Order.GetAllOrdersController)
		orderRoutes.GET("/:id", ctx.Ctl.Order.GetOrderByIDController)
		orderRoutes.PUT("/update/:id", ctx.Ctl.Order.UpdateOrderStatusController)
		orderRoutes.DELETE("/:id", ctx.Ctl.Order.DeleteOrderController)
		orderRoutes.GET("/detail/:id", ctx.Ctl.Order.GetOrderDetailController)
	}

	return r
}

func authRoutes(r *gin.Engine, ctx infra.ServiceContext) {
	authGroup := r.Group("/auth")
	authGroup.GET("/login", ctx.Ctl.User.LoginController)
	authGroup.GET("/check-email", ctx.Ctl.User.CheckEmailUserController)
	authGroup.POST("/register", ctx.Ctl.User.CreateUserController)
	authGroup.PATCH("/reset-password", ctx.Ctl.User.ResetUserPasswordController)
}
