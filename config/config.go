package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	InfluxDB struct {
		URL    string
		Token  string
		Org    string
		Bucket string
	}
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("警告: .env 文件加载失败: %v", err)
	}

	cfg := &Config{}
	cfg.InfluxDB.URL = getEnv("INFLUXDB_URL", "http://localhost:8086")
	cfg.InfluxDB.Token = getEnv("INFLUXDB_TOKEN", "")
	cfg.InfluxDB.Org = getEnv("INFLUXDB_ORG", "org")
	cfg.InfluxDB.Bucket = getEnv("INFLUXDB_BUCKET", "bucket")
	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
