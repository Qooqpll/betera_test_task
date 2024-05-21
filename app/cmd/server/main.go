package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	adapter "github.com/qooqpll/betera_test"
	"github.com/qooqpll/betera_test/internal/controller"
	"github.com/qooqpll/betera_test/internal/repository"
	"github.com/qooqpll/betera_test/internal/service"
	"log"
)

func main() {
	err := godotenv.Load()
	router := gin.Default()
	db := adapter.ConnectWithDB()

	apodRepo := repository.NewApodRepo(db)
	apodService := service.NewApodService(apodRepo)
	apodController := controller.NewApodController(apodService)

	router.GET("/apod", apodController.All)

	if err != nil {
		log.Println(err)
	}

	router.Run()
}
