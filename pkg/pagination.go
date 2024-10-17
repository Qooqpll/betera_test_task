package pkg

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Size       int    `json:"perPage"`
	Page       int    `json:"page"`
	SortField  string `json:"sortBy"`
	Direction  string `json:"sortDirection"`
	TotalRows  int64  `json:"totalElements"`
	TotalPages int    `json:"totalPages"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetSize()
}

func (p *Pagination) GetSize() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSortField() string {
	if p.SortField == "" {
		p.SortField = "created_at"
	}
	return p.SortField
}

func (p *Pagination) GetDirection() string {
	if p.Direction == "" {
		p.Direction = "asc"
	}
	return p.Direction
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetSize())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetSize()).Order(pagination.GetSortField() + " " + pagination.GetDirection())
	}
}

func (p *Pagination) Bind(c *gin.Context) (err error) {
	strSize := c.Query("perPage")
	strPage := c.Query("page")
	sortField := c.Query("sortBy")
	direction := c.Query("sortDirection")
	var size int
	var page int

	if len(strSize) > 0 {
		size, err = strconv.Atoi(strSize)
		if err != nil {
			return
		}
	}

	if len(strPage) > 0 {
		page, err = strconv.Atoi(strPage)
		if err != nil {
			return
		}
	}
	p.Size = size
	p.Page = page
	p.SortField = sortField
	p.Direction = direction
	return
}
