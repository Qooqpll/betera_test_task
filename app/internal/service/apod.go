package service

import (
	"fmt"
	"github.com/qooqpll/betera_test/internal/model"
	"github.com/qooqpll/betera_test/internal/repository"
	"log"
)

type ApodService interface {
	All(date string) []model.Apod
	DailyTask()
}

type apodService struct {
	apodRepo repository.ApodRepo
}

func NewApodService(apodRepo repository.ApodRepo) *apodService {
	return &apodService{apodRepo: apodRepo}
}

func (a *apodService) All(date string) []model.Apod {
	apods := []model.Apod{}
	if len(date) == 0 {
		return a.apodRepo.All()
	}

	apod, err := a.apodRepo.GetByDate(date)
	if err != nil {
		log.Println(err)
	}

	return append(apods, apod)
}

func (a *apodService) DailyTask() {
	fmt.Println("cron is started")
	response := SendApodRequest()
	SaveImage(response.Hdurl)
	a.apodRepo.Insert(response)
}
