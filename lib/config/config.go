package config

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"jirago/lib/logger"
)

var cfgFile string

func Setup() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Error("home dir failed", logger.Fields{"err": err})
			os.Exit(1)
		}

		// Search config in home directory with name ".jirago" (without extension).
		viper.AddConfigPath(home + "/.config")
		viper.SetConfigName("jirago")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error("no config found", logger.Fields{"file": viper.ConfigFileUsed()})
		}
	}
}
