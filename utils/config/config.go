// Package config is reserved for all config initializing functions.
//
package config

import (
	"strings"

	"github.com/spf13/viper"
)

// ReadConfig sets default Config-Values and reads the Config-File
//
func ReadConfig() (bool, error) {
	configSetDefaults()

	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/hanson/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.hanson") // call multiple times to add many search paths
	viper.AddConfigPath(".")             // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return false, err
	}

	return true, nil
}

// GetStringValue returns String-Value of a key from config
//
// PARAMETERS:
//    key - Config-Key to read out
//
// RETURN:
//    String of Key-Value
//
func GetStringValue(key string) string {
	return viper.GetString(key)
}

// GetBoolValue returns Bool-Value of a key from config. If value is "true",
// then returns a bool true
//
// PARAMETERS:
//    key - Config-Key to read out
//
// RETURN:
//    Bool of Key-Value
//
func GetBoolValue(key string) bool {
	if strings.ToLower(viper.GetString(key)) == "true" {
		return true
	}

	return false
}

// GetNumericValue returns int Value of a key from config.
//
// PARAMETERS:
//    key - Config-Key to read out
//
// RETURN:
//    int of Key-Value
//
func getNumericValue(key string) int {
	return viper.GetInt(key)
}

// configSetDefaults sets default config-values
//
func configSetDefaults() {
	// default configuration, if some entries are missing
	viper.SetDefault("LoggingFileUse", "false")
	viper.SetDefault("LoggingFile", "log/hanson.log")
	viper.SetDefault("LoggingLevel", "warning")
	viper.SetDefault("WebserverDebug", "false")
	viper.SetDefault("WebserverPort", "8080")

	// Config-Parameter only for testing
	viper.SetDefault("TestingString", "yes")
	viper.SetDefault("TestingBool", true)
}
