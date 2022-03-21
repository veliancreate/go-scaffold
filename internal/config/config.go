package config

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type Config struct {
	Server ServerConfig
	DB     DBConfig
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port: "8000",
		},
		DB: DBConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "postgres",
			DbName:   "book_store",
		},
	}
}
