package controller

import (
	"fmt"
	"net/http"
	"strconv"

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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to get coupon: %w", err).Error()})
		return
	}
	c.JSON(http.StatusOK, coupon)

}

func (ac *AdminController) CreateCategory(c *gin.Context) {
	var category db.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input category format"})
		return
	}

	err := ac.Repositories.CategoriesRepository.CreateCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to create category: %w", err).Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (ac *AdminController) GetCategoryIDByName(c *gin.Context) {
	categoryName := c.Param("categoryName")
	categoryID, err := ac.Repositories.CategoriesRepository.GetCategoryIDByName(c.Request.Context(), categoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to get category ID: %w", err).Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categoryID": categoryID})
}

func (ac *AdminController) CreateMedicineWithCategory(c *gin.Context) {
	var medicine db.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input medicine format"})
		return
	}
	err := ac.Repositories.MedicineRepositoryService.CreateMedicine(c.Request.Context(), &medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to create medicine: %w", err).Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Medicine created successfully"})
}

func (ac *AdminController) UpdateMedicine(c *gin.Context) {
	var medicine db.Medicine
	if err := c.ShouldBindJSON(&medicine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input medicine format"})
		return
	}

	err := ac.Repositories.MedicineRepositoryService.UpdateMedicine(c.Request.Context(), &medicine)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to update medicine: %w", err).Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Medicine updated successfully"})
}

func (ac *AdminController) DeleteMedicine(c *gin.Context) {
	medicineID := c.Param("medicineID")
	medicineIDInt, err := strconv.Atoi(medicineID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medicine ID format"})
		return
	}

	err = ac.Repositories.MedicineRepositoryService.DeleteMedicine(c.Request.Context(), medicineIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to delete medicine: %w", err).Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Medicine deleted successfully"})
}

func (ac *AdminController) GetAllMedicines(c *gin.Context) {
	medicines, err := ac.Repositories.MedicineRepositoryService.GetAllMedicines(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Errorf("failed to get all medicines: %w", err).Error()})
		return
	}
	c.JSON(http.StatusOK, medicines)
}