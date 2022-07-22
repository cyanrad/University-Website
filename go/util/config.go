package util

import "github.com/spf13/viper"

// >> App confgurations
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// >> reading config from env file/variables
func LoadConfig(path string) (config Config, err error) {
	// >> env file path
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// >> Read environment variable
	viper.AutomaticEnv()

	// >> reading configurtions
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// >> storing configurations
	err = viper.Unmarshal(&config)
	return
}
