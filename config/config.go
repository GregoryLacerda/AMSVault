package config

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	MongoURI          string `mapstructure:"MONGO_URI"`
	MongoDB           string `mapstructure:"MONGO_DB"`
	MongoCollection   string `mapstructure:"MONGO_COLLECTION"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret         string `mapstructure:"JWT_SECRET"`
	JWTExpirationTime int    `mapstructure:"JWT_EXPIRATION_TIME"`
	TokenAuth         *jwtauth.JWTAuth
	MAL_API_URL       string `mapstructure:"MAL_API_URL"`
	MAL_API_AUTH_URL  string `mapstructure:"MAL_API_AUTH_URL"`
	MAL_CLIENT_ID     string `mapstructure:"MAL_CLIENT_ID"`
	MAL_CLIENT_SECRET string `mapstructure:"MAL_CLIENT_SECRET"`
	MAL_GRANT_TYPE    string `mapstructure:"MAL_GRANT_TYPE"`
	MAL_REFRESH_TOKEN string `mapstructure:"MAL_REFRESH"`
	MAL_TOKEN         string `mapstructure:"MAL_TOKEN"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.TokenAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)
	return &config, nil
}
