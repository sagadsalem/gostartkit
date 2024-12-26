package app

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/sajadsalem/startkit/config"
	"github.com/sajadsalem/startkit/db"
	"github.com/sajadsalem/startkit/pkg/validation"
	"gorm.io/gorm"
)

type App struct {
	Validator *validator.Validate
	Addr      string
	DB        *gorm.DB
	Configs   *config.Configs
}

var err error

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewApp(addr string) (*App, error) {
	configs := config.LoadConfigs()

	db, err := db.ConnectAndMigrate(configs.DB)
	if err != nil {
		return nil, err
	}

	v := validator.New()
	validation.RegisterCustomValidations(v)

	return &App{
		Validator: v,
		Addr:      addr,
		DB:        db,
		Configs:   configs,
	}, nil
}
