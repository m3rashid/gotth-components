package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m3rashid/gotth-components/components"
)

const componentsJsonFileName = "components.json"

func generateJsonConfig() {
	jsonData, err := json.MarshalIndent(components.ConfigWithDependencies, "", "  ")
	if err != nil {
		fmt.Println("Could not convert file to json: ", err)
		return
	}

	file, err := os.Create(componentsJsonFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data successfully written to file")
}
