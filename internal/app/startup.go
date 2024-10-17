package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/madflojo/tasks"
	log "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	configApp "golang-test-task-betera/internal/config"
	"golang-test-task-betera/internal/repository"
	"time"
)

func Run() {
	router := gin.Default()
	router.Use(cors.Default())
	setupRoutes(router)
	router.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerFiles.Handler)(context)
	})

	config := configApp.GetConfigurationInstance()

	scheduler := tasks.New()
	defer scheduler.Stop()
	_, err := scheduler.Add(&tasks.Task{
		Interval: 24 * time.Hour,
		TaskFunc: func() error {
			err := repository.SaveDailyApodInfo()
			return err
		},
		ErrFunc: func(e error) {
			log.Printf("an error occurred when executing task %s", e)
		},
	})
	if err != nil {
		log.Warningf("error executing task: %s", err)
	}

	err = router.Run(config.AppPort)
	if err != nil {
		log.Fatalf("couldn't run app on port: " + config.AppPort)
	}
}

func setupRoutes(router *gin.Engine) {
	repository.SetupRoutes(router)
}
