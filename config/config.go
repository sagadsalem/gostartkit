package config

type Configs struct {
	DB   *DatabaseConfig
	Mail *MailConfig
}

func LoadConfigs() *Configs {
	return &Configs{
		DB:   LoadDatabaseConfig(),
		Mail: LoadMailConfig(),
	}
}
