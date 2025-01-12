package configs

var cfg *Config

type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

type JWTConfig struct {
	SecretKey  string
	ExpireTime int
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
}

func LoadConfig(parh string) (*Config, error) {
	// Load config from file

	return nil, nil
}
