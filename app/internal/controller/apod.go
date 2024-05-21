package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/qooqpll/betera_test/internal/service"
	"net/http"
)

type ApodController interface {
	All(context *gin.Context)
}

type apodController struct {
	apodService service.ApodService
}

func NewApodController(apod service.ApodService) *apodController {
	return &apodController{apodService: apod}
}

func (a *apodController) All(context *gin.Context) {
	date := context.Query("date")
	response := a.apodService.All(date)
	context.JSON(
		http.StatusOK,
		gin.H{"response": response},
	)
}
