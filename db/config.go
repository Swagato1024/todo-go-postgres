package db

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func getConfig() Config {
	return Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "secret",
		DBName:   "todo",
		SslMode:  "disabled",
	}
}
