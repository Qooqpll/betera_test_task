package adapter

import (
	"fmt"
	"github.com/qooqpll/betera_test/config"
	"github.com/qooqpll/betera_test/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectWithDB() *gorm.DB {
	dbConfig := config.GetDBConfig()
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Dbname,
		dbConfig.Port,
		dbConfig.Sslmode,
		dbConfig.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {

	}
	db.AutoMigrate(model.Apod{})
	return db
}
