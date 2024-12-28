package config

type Config struct {
	Port   string
	JWTKey []byte
}

func NewConfig() *Config {
	return &Config{
		Port:   ":8080",
		JWTKey: []byte("your_secret_key"), // In production, load from environment
	}
}
