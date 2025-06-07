package controller

import (
	"github.com/AjayKodavati/CMS/repository"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	Repositories repository.DBRepositories
}

func (ac *AdminController) CreateCoupon(c *gin.Context) {

}

func (ac *AdminController) CreateCategory(c *gin.Context) {
	
}

func (ac *AdminController) CreateMedicineWithCategory(c *gin.Context) {

}