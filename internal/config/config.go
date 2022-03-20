package config

type ServerConfig struct {
	Port string
}

type DBConfig struct {
	URL string
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
			URL: "postgres://postgres:postgres@localhost:5432/book_store",
		},
	}
}
