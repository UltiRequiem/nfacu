package cmd

import (
	"fmt"
	"github.com/UltiRequiem/nfacu/internal"
)

// Start the process
func Main() {
	configPath := getArguments()

	fmt.Println("Logging enabled.")

	projectsConfig, errorGettingConfig := getConfig(configPath)

	if errorGettingConfig != nil {
		fmt.Printf("Error getting config: %s\n", errorGettingConfig.Error())
		return
	}

	fmt.Printf(`Config read from "%s".`+ "\n\n", configPath)

	for _, project := range projectsConfig {
		rawProjectConfig, errorGettingProjectConfig := getProjectConfig(project.Path)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error getting project config: %s\n", errorGettingProjectConfig.Error())
			return
		}

		configRawData := ""

		for _, line := range rawProjectConfig {
			lineAdded := false
			for key := range project.Settings {
				if containsInString(line, key) {
					configRawData += internal.ParseLine(key, project.Settings[key])
					lineAdded = true
					fmt.Printf(`Changing "%s" property to "%s" on "%s".`+"\n", key, project.Settings[key], project.Path)
				}
			}

			if !lineAdded {
				configRawData += line
			}

			configRawData += "\n"
		}

		errorSavingAppConfig := saveConfigFile(project.Path, configRawData[:len(configRawData)-1])

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error saving project config: %s\n", errorSavingAppConfig.Error())
			return
		}

		fmt.Printf("\n%s updated successfully!\n", project.Path)
	}
}
