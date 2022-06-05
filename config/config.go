package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	DataDir string `yaml:"data_dir,omitempty"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		path := "./data/config"
		viper.SetConfigName("config")        // config file name without extension
		viper.SetConfigType("yaml")          // yaml type
		viper.AddConfigPath("./data/config") // config file path
		viper.AutomaticEnv()                 // read value ENV variable

		// Set default value
		viper.SetDefault("DataDir", "data")

		// write config if not present
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			panic(fmt.Sprintln("fatal error creating config file directory \n", err))
		}
		viper.SafeWriteConfig()

		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Sprintln("fatal error config file: default \n", err))
		}

		viper.Unmarshal(&config)

		// Config
		log.Debug().Msgf("DataDir : %s", config.DataDir)
	}
	return config
}

func init() {

}
