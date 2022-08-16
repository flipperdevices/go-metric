package models

type Config struct {
	DBAddr string `env:"DB_ADDR" envDefault:"localhost:9000"`
	DBName string `env:"DB_NAME" envDefault:"metric"`
}
