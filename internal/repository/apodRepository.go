package repository

import (
	"github.com/gin-gonic/gin"
	"golang-test-task-betera/internal/config"
	"golang-test-task-betera/internal/model"
	"golang-test-task-betera/internal/status"
	"golang-test-task-betera/pkg"
	"net/http"
	"time"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/apod")
	api.GET("/get", GetAstrologers)
	api.GET("/get-by-date", GetAstrologersByDate)
}

func GetAstrologers(c *gin.Context) {
	var pr model.PageableResponse
	var p pkg.Pagination
	err := p.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.PaginationDataValidationError()))
		return
	}

	var apod []model.Apod
	paginateFunc := pkg.Paginate(&model.Apod{}, &p, config.GetDBInstance())
	err = paginateFunc(config.GetDBInstance()).Find(&apod).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, pr.ErrorResponse(status.ExecutingQueryError("find all")))
		return
	}

	config.GetDBInstance().Create(&model.Apod{})

	c.JSON(http.StatusOK, pr.New(apod, p, nil))
}

func GetAstrologersByDate(c *gin.Context) {
	var pr model.PageableResponse
	var p pkg.Pagination
	err := p.Bind(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.PaginationDataValidationError()))
		return
	}

	d := c.Query("date")
	date, err := time.Parse("2006-01-02", d)
	if err != nil {
		c.JSON(http.StatusBadRequest, pr.ErrorResponse(status.DataConversionError()))
		return
	}

	var apod []model.Apod
	paginateFunc := pkg.Paginate(&model.Apod{}, &p, config.GetDBInstance())
	err = paginateFunc(config.GetDBInstance()).Where("day_info.apod.date = ?", date).Find(&apod).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, pr.ErrorResponse(status.ExecutingQueryError("find by date")))
		return
	}

	c.JSON(http.StatusOK, pr.New(apod, p, nil))
}
