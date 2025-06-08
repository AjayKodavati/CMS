package controller

import (
	"fmt"
	"net/http"

	"github.com/AjayKodavati/CMS/db"
	"github.com/AjayKodavati/CMS/repository"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Repositories repository.DBRepositories
}

func (ac *AdminController) CreateCoupon(c *gin.Context) {
var coupon db.Coupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input coupon format"})
		return
	}

	err := ac.Repositories.CouponRepository.CreateCoupon(c.Request.Context(), &coupon)
	if err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to create coupon: %w", err).Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Coupon created successfully"})
}

func (ac *AdminController) DeleteCoupon(c *gin.Context) {
	couponCode := c.Param("couponCode")

	if err := ac.Repositories.CouponRepository.DeleteCoupon(c.Request.Context(), couponCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon deleted successfully"})
}

func (ac *AdminController) UpdateCoupon(c *gin.Context) {
	var coupon db.Coupon
	if err := c.ShouldBindBodyWithJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input coupon format"})
		return
	}

	if err := ac.Repositories.CouponRepository.UpdateCoupon(c.Request.Context(), &coupon); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "coupon updated successfully"})
}

func (ac *AdminController) GetCouponByID(c *gin.Context) {
	couponCode := c.Param("couponCode")

	coupon, err := ac.Repositories.CouponRepository.GetCouponByID(c.Request.Context(), couponCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve coupon"})
		return
	}
	c.JSON(http.StatusOK, coupon)

}

func (ac *AdminController) CreateCategory(c *gin.Context) {
	
}

func (ac *AdminController) CreateMedicineWithCategory(c *gin.Context) {

}
