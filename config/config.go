package config

type Config struct {
	InfluxDB struct {
		URL    string
		Token  string
		Org    string
		Bucket string
	}
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.InfluxDB.URL = "http://localhost:8086"
	cfg.InfluxDB.Token = "your-token"
	cfg.InfluxDB.Org = "your-org"
	cfg.InfluxDB.Bucket = "your-bucket"
	return cfg
}
