package main

import (
	"fmt"
	adapter "github.com/qooqpll/betera_test"
	"github.com/qooqpll/betera_test/internal/repository"
	"github.com/qooqpll/betera_test/internal/service"
	"github.com/robfig/cron/v3"
)

func main() {
	db := adapter.ConnectWithDB()
	apodRepo := repository.NewApodRepo(db)
	apodService := service.NewApodService(apodRepo)

	c := cron.New()

	_, err := c.AddFunc("@daily", apodService.DailyTask)
	if err != nil {
		fmt.Println("Error scheduling task:", err)
		return
	}

	c.Start()
	select {}
}
