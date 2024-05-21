package repository

import (
	"errors"
	"github.com/qooqpll/betera_test/internal/model"
	"gorm.io/gorm"
)

type ApodRepo interface {
	Insert(apod model.Apod) model.Apod
	All() []model.Apod
	GetByDate(date string) (model.Apod, error)
}

type apodRepo struct {
	db *gorm.DB
}

func NewApodRepo(db *gorm.DB) *apodRepo {
	return &apodRepo{db: db}
}

func (r *apodRepo) Insert(apod model.Apod) model.Apod {
	r.db.Create(&apod)
	r.db.First(&apod, apod.ID)
	return apod
}

func (repo *apodRepo) All() []model.Apod {
	apod := []model.Apod{}
	repo.db.Order("id desc").Find(&apod)
	return apod
}

func (repo *apodRepo) GetByDate(date string) (model.Apod, error) {
	apod := model.Apod{}
	err := repo.db.First(&apod, "date = ?", date).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return apod, err
	}
	return apod, nil
}
