package model

type Apod struct {
	ID          uint   `gorm:"primary_key:auto_increment"`
	Date        string `gorm:"type:varchar(100)" json:"date"`
	Explanation string `gorm:"type:text" json:"explanation"`
	Hdurl       string `gorm:"type:varchar(255)" json:"hdurl"`
	MediaType   string `gorm:"type:varchar(100)" json:"media_type"`
	Title       string `gorm:"type:text" json:"title"`
	Url         string `gorm:"type:varchar(255)" json:"url"`
}
