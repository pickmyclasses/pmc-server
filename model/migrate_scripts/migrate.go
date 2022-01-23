package main

import (
	"fmt"
	"log"
	"os"

	"pmc_server/config"
	model "pmc_server/model"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.username"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.dbname"),
		viper.GetInt("postgres.port"))
	stdLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: stdLogger,
	})

	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.Review{})
}
