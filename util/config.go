package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string ) (config Config , err error ) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	
	 
	viper.SetConfigName("app")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("⚠️ No app.env file found, trying local.app.env...")
		// If app.env is missing, try local.app.env
		viper.SetConfigName("local.app")
		if localErr := viper.ReadInConfig(); localErr != nil {
			fmt.Println("⚠️ No local.app.env found, relying on environment variables only.")
		}
	}

	err = viper.Unmarshal(&config)
	return

}