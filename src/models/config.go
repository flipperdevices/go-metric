package models

type Config struct {
	DBAddr string `env:"DB_ADDR" envDefault:"localhost:19000"`
	DBName string `env:"DB_NAME" envDefault:"metric"`
}
