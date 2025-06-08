package router

import (
	"github.com/AjayKodavati/CMS/controller"
	"github.com/AjayKodavati/CMS/server"
)

func SetUpAdminRoutes(server *server.Server) {
	adminGroup := server.Router.Group("/admin")
	adminController := &controller.AdminController{
		Repositories: server.RepositoryService,
	}

	couponGroup := adminGroup.Group("/coupon")
	{
		couponGroup.DELETE("/delete/:couponCode", adminController.DeleteCoupon)
		couponGroup.PUT("/update", adminController.UpdateCoupon)
		couponGroup.GET("/:couponID", adminController.GetCouponByID)
		couponGroup.POST("/create", adminController.CreateCoupon)
	}

	// categoryGroup := adminGroup.Group("/category")
	// {
	// 	categoryGroup.POST("/create", adminController.CreateCategory)	
	// }

	// medicineGroup := adminGroup.Group("/medicine")
	// {
	// 	medicineGroup.POST("/create", adminController.CreateMedicineWithCategory)
	// 	medicineGroup.PUT("/update/:medicineID", adminController.UpdateMedicine)
	// 	medicineGroup.DELETE("/delete/:medicineID", adminController.DeleteMedicine)
	// }
}