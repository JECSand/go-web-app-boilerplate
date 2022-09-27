package cmd

import (
	"encoding/json"
	"os"
)

// Configuration is a struct designed to hold the applications variable configuration settings
type Configuration struct {
	Port           string
	APIHost        string
	SessionManager string
	RedisHost      string
	RedisPassword  string
	ENV            string
}

// ConfigurationSettings is a function that reads a json configuration file and outputs a Configuration struct
func ConfigurationSettings(env string) (*Configuration, error) {
	confFile := "conf.json"
	if env == "test" {
		confFile = "test_conf.json"
	}
	file, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	configurationSettings := Configuration{}
	err = decoder.Decode(&configurationSettings)
	if err != nil {
		return nil, err
	}
	return &configurationSettings, nil
}

// InitializeEnvironment sets environmental variables
func (c *Configuration) InitializeEnvironment() {
	os.Setenv("PORT", c.Port)
	os.Setenv("API_HOST", c.APIHost)
	os.Setenv("SESSION_MANAGER", c.SessionManager)
	os.Setenv("REDIS_HOST", c.RedisHost)
	os.Setenv("REDIS_PASSWORD", c.RedisPassword)
	os.Setenv("ENV", c.ENV)
}
