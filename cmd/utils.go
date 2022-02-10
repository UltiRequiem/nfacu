package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/UltiRequiem/nfacu/internal"
)

func getConfig(configFilePath string) (internal.NFACUConfig, error) {
	configData, errorReadingConfig := os.ReadFile(configFilePath)

	if errorReadingConfig != nil {
		return nil, errorReadingConfig
	}

	var projectsConfig internal.NFACUConfig

	errorUnmarshalling := json.Unmarshal(configData, &projectsConfig)

	if errorUnmarshalling != nil {
		return nil, errorUnmarshalling
	}

	return projectsConfig, nil
}

func getProjectConfig(path string) ([]string, error) {
	projectConfig, errorReadingProjectConfig := os.ReadFile(path)

	if errorReadingProjectConfig != nil {
		return nil, errorReadingProjectConfig
	}

	return strings.Split(string(projectConfig), "\n"), nil
}

func saveConfigFile(path, parsedConfig string) error {
	errorWritingProjectConfig := os.WriteFile(path, []byte(parsedConfig), 0644)

	if errorWritingProjectConfig != nil {
		return errorWritingProjectConfig
	}

	return nil
}

func getArguments() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}

	return "nfacu.json"
}

func containsInString(text, textToSearch string) bool {
	return strings.Contains(text, fmt.Sprintf(`"%s"`, textToSearch))
}
