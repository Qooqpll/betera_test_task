package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	directory = "images"
)

func SaveImage(hdurl string) {
	responseImage, _ := http.Get(hdurl)
	defer responseImage.Body.Close()

	file := createDirectoryAndFile()

	defer file.Close()

	_, err := io.Copy(file, responseImage.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func createDirectoryAndFile() *os.File {
	year, month, day := time.Now().Date()
	y := strconv.Itoa(year)
	m := strconv.Itoa(int(month))
	d := strconv.Itoa(day)
	dir := fmt.Sprintf("%s/%s/%s/%s", directory, y, m, d)
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(dir, 0700)
		}
	}

	s := strconv.Itoa(int(time.Now().Unix()))
	file, err := os.Create(dir + "/" + s + ".jpg")
	if err != nil {
		log.Fatal(err)
	}

	return file
}
