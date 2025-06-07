package router

import (
	"github.com/AjayKodavati/CMS/server"
)

func SetUpCouponRoutes(server *server.Server) {
	couponGroup := server.Router.Group("/coupons")
	couponController := &controller.CouponController{
		Repositories: server.RepositoryService,
	}

	couponGroup.POST("/applicable", couponController.ApplicableCoupons)
	couponGroup.POST("/validate", couponController.ValidateCoupon)
	
}