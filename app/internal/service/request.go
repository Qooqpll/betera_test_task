package service

import (
	"bytes"
	"encoding/json"
	"github.com/qooqpll/betera_test/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	UrlApod = "https://api.nasa.gov/planetary/apod"
)

func SendApodRequest() model.Apod {
	apiKey := os.Getenv("API_KEY")
	resBody := makeRequest(http.MethodGet, UrlApod, "api_key="+apiKey)
	var response model.Apod
	err := json.Unmarshal(resBody, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func makeRequest(method, url, params string) []byte {
	req, err := http.NewRequest(method, url+"?"+params, bytes.NewBuffer([]byte{}))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	return body
}
