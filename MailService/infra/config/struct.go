package config

type Smtp struct {
	Host     string
	Port     int
	Username string
	Password string
}

type Config struct {
	Smtp Smtp
}
