package controllers

import (
	"mnc_test/models"
	"mnc_test/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customerController struct {
	router          *gin.RouterGroup
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerController(router *gin.RouterGroup, customerUsecase usecase.CustomerUsecase) *customerController {
	controller := customerController{
		router:          router,
		customerUsecase: customerUsecase,
	}

	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)

	return &controller
}

func (c *customerController) Create(ctx *gin.Context) {
	var customer models.Customer

	if err := ctx.BindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if customer.Name == "" || customer.Email == "" || customer.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "one or more field are missing",
		})
		return
	}

	res, err := c.customerUsecase.Create(&customer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *customerController) Login(ctx *gin.Context) {
	var customer models.Customer

	if err := ctx.BindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if customer.Email == "" || customer.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "one or more field are missing",
		})
		return
	}

	token, err := c.customerUsecase.Login(customer.Email, customer.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
