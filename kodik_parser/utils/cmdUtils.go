package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func InitLogger() error {
	file, err := os.OpenFile("kodikParser.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.New("ошибка при создании лога")
	}

	log.SetOutput(file)

	return nil
}

func OpenInMpvNet(result HandleResult, config *Config) error {
	var commands []string

	commands = append(commands, "append")

	if config.DownloadResults {
		for _, res := range result.Results {
			commands = append(commands, res.Path)
		}
	} else {
		for _, res := range result.Results {
			commands = append(commands, res.Video)
		}
	}

	cmd := exec.Command(config.MpvNetExecutable, commands...)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error while opening with MVP.NET!")
		return err
	}

	return nil
}

func PrintResults(result HandleResult) {
	for _, res := range result.Results {
		fmt.Printf("Серия %s: %s\n", res.Seria.Num, res.Video)
	}
}
