package model

import (
	"gorm.io/gorm"
)

type ApodResponse struct {
	Copyright      string     `json:"copyright"`
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	HdUrl          string     `json:"hdurl"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	Url            string     `json:"url"`
}

type Apod struct {
	gorm.Model
	Copyright      string     `json:"copyright"`
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	Image          []byte     `json:"image"`
}

func (Apod) TableName() string {
	return "day_info.apod"
}

func (ar *ApodResponse) ToApod(imgBytes []byte) *Apod {
	return &Apod{
		Copyright:      ar.Copyright,
		Date:           ar.Date,
		Explanation:    ar.Explanation,
		MediaType:      ar.MediaType,
		ServiceVersion: ar.ServiceVersion,
		Title:          ar.Title,
		Image:          imgBytes,
	}
}
