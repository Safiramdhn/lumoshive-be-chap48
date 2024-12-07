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
	r.GET("/api/banner", ctx.Ctl.Banner.GetBannerByIDController)
	r.DELETE("/api/banner", ctx.Ctl.Banner.DeleteBannerController)
	r.PUT("/api/banner", ctx.Ctl.Banner.UpdateBannerController)
	r.POST("/api/create-banner", ctx.Ctl.Banner.CreateBannerController)
	return r
}
