package config

import (
	"os"
	"strconv"
)

type MailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func LoadMailConfig() *MailConfig {
	p, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		p = 587
	}

	return &MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     p,
		Username: os.Getenv("MAIL_USERNAME"),
		Password: os.Getenv("MAIL_PASSWORD"),
	}
}
