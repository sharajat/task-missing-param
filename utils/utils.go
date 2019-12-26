package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"task-missing-param/structure"
)

// ExtractBootDetails extracts json data of boot
func ExtractBootDetails() structure.Boot {

	jsonFile, err := os.Open("introspectedData.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonString, _ := ioutil.ReadAll(jsonFile)
	result := structure.Data{}
	json.Unmarshal([]byte(jsonString), &result)
	return result.HardwareDetails.Boot

}

// ExtractCPUDetails extracts json data of cpu
func ExtractCPUDetails() structure.CPU {

	jsonFile, err := os.Open("introspectedData.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	jsonString, _ := ioutil.ReadAll(jsonFile)
	result := structure.Data{}
	json.Unmarshal([]byte(jsonString), &result)
	return result.HardwareDetails.CPU

}
