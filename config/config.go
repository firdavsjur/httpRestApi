package config

type Config struct {
	ServerHost string
	ServerPort string

	PostgresHost     string
	PostgresUser     string
	PostgresDatabase string
	PostgresPassword string
	PostgresPort     string
}

func Load() Config {
	cfg := Config{}

	cfg.ServerHost = "localhost"
	cfg.ServerPort = ":4000"

	cfg.PostgresHost = "localhost"
	cfg.PostgresUser = "postgres"
	cfg.PostgresDatabase = "shopcart"
	cfg.PostgresPassword = "0000"
	cfg.PostgresPort = "5432"

	return cfg
}
