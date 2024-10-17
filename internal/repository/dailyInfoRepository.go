package repository

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	configApp "golang-test-task-betera/internal/config"
	"golang-test-task-betera/internal/model"
	"io"
	"net/http"
)

func SaveDailyApodInfo() error {
	apodResponse, err := getDayInfo()
	if err != nil {
		return err
	}
	imgBytes, err := getImage(apodResponse.Url)
	if err != nil {
		return err
	}

	apod := apodResponse.ToApod(imgBytes) //todo rename to setImage
	err = configApp.GetDBInstance().Create(&apod).Error
	if err != nil {
		return err
	}

	log.Info("[integration] day info saved")
	return nil
}

func getDayInfo() (model.ApodResponse, error) {
	var ar model.ApodResponse

	config := configApp.GetConfigurationInstance()
	url := config.Url + config.ApiKey

	resp, err := http.Get(url)
	fmt.Println(resp.Body)
	if err != nil {
		return ar, fmt.Errorf("failed to make HTTP request: %w", err)
	}

	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("failed to close response body: %v", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return ar, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	fmt.Println(ar)
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		return ar, fmt.Errorf("failed to decode response: %w", err)
	}

	return ar, nil
}

func getImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			log.Printf("failed to close response body: %v", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	imgBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return imgBytes, nil
}
