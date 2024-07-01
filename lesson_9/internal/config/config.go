package config

import (
	"brave_signer/internal/logger"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ContextKey uint

const ViperKey ContextKey = 0

// LoadYamlConfig loads configuration from a YAML file
func LoadYamlConfig() (*viper.Viper, error) {
	localViper := viper.New()
	localViper.SetConfigName("config")
	localViper.SetConfigType("yaml")
	localViper.AddConfigPath(".")

	err := localViper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Info("Config file not found, using CLI params or default settings...")
		} else {
			return localViper, fmt.Errorf("found config file, but encountered an error: %v", err)
		}
	}
	return localViper, nil
}

func BindFlags(cmd *cobra.Command, v *viper.Viper) error {
	var firstErr error

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if err := v.BindPFlag(flag.Name, flag); err != nil {
			if firstErr == nil { // Store the first error encountered
				firstErr = fmt.Errorf("error binding flag '%s': %v", flag.Name, err)
			}
			logger.Warn(err)
		}

		if !flag.Changed && v.IsSet(flag.Name) {
			if err := cmd.Flags().Set(flag.Name, v.GetString(flag.Name)); err != nil {
				if firstErr == nil { // Store the first error encountered
					firstErr = fmt.Errorf("error setting flag '%s' from config: %v", flag.Name, err)
				}
				logger.Warn(err)
			}
		}
	})

	return firstErr
}
