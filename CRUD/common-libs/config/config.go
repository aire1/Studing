package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	PostgresUrl = getConfig().PosgresUrl
)

type config struct {
	PosgresUrl string
}

func ParseConfig() (config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return config{}, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return config{}, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	config := config{
		PosgresUrl: result["databaseUrl"].(string),
	}

	return config, nil
}

func getConfig() config {
	conf, err := ParseConfig()
	if err != nil {
		log.Fatalf("Fatal! %s", err.Error())
	}

	return conf
}
