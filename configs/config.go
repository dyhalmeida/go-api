package configs

import (
	"os"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *config

type databaseConfig struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	name     string
}

type serverConfig struct {
	port string
}

type jwtConfig struct {
	secretKey string
	expiresIn int
	tokenAuth *jwtauth.JWTAuth
}

type config struct {
	database databaseConfig
	server   serverConfig
	jwt      jwtConfig
}

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.SetConfigFile(workDir + "/.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	cfg = &config{
		database: databaseConfig{
			driver:   viper.GetString("db_driver"),
			host:     viper.GetString("db_host"),
			port:     viper.GetString("db_port"),
			user:     viper.GetString("db_user"),
			password: viper.GetString("db_password"),
			name:     viper.GetString("db_name"),
		},
		server: serverConfig{
			port: viper.GetString("server_port"),
		},
		jwt: jwtConfig{
			secretKey: viper.GetString("jwt_secret_key"),
			expiresIn: viper.GetInt("jwt_expires_in"),
		},
	}
	cfg.jwt.tokenAuth = jwtauth.New("HS256", []byte(cfg.jwt.secretKey), nil)
}

func NewConfig() *config {
	return cfg
}

func (c *config) GetDatabaseDriver() string {
	return c.database.driver
}

func (c *config) GetDatabaseHost() string {
	return c.database.host
}

func (c *config) GetDatabasePort() string {
	return c.database.port
}

func (c *config) GetDatabaseUser() string {
	return c.database.user
}

func (c *config) GetDatabasePassword() string {
	return c.database.password
}

func (c *config) GetDatabaseName() string {
	return c.database.name
}

func (c *config) GetJwtSecretKey() string {
	return c.jwt.secretKey
}

func (c *config) GetJwtExpiresIn() int {
	return c.jwt.expiresIn
}

func (c *config) GetTokenAuth() *jwtauth.JWTAuth {
	return c.jwt.tokenAuth
}

func (c *config) GetServerPort() string {
	return c.server.port
}
