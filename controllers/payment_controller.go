package controllers

import (
	"mnc_test/models"
	"mnc_test/usecase"
	"mnc_test/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type paymentController struct {
	router         *gin.RouterGroup
	paymentUsecase usecase.PaymentUsecase
}

func NewPaymentUsecase(router *gin.RouterGroup, paymentUsecase usecase.PaymentUsecase) *paymentController {
	controller := paymentController{
		router:         router,
		paymentUsecase: paymentUsecase,
	}

	router.Use(utils.AuthMiddleware())
	router.POST("/payment", controller.Create)

	return &controller
}

func (c *paymentController) Create(ctx *gin.Context) {
	var payment models.Payment

	if err := ctx.BindJSON(&payment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	if payment.Customer_Id == 0 || payment.Bill == 0 || payment.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "one or more field are missing",
		})
		return
	}

	res, err := c.paymentUsecase.Create(&payment)
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
