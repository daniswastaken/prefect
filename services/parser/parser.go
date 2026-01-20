package parser

import (
	"encoding/json"
	"os"
	"log"
	"prefect/services/sys"
)

type SysData struct {
	CPUCores int     `json:"cpu_cores"`
	CPUUsage int `json:"cpu_usage"`
	// Add other system data fields as needed
}

func SysDataParser() {
	// Data Structures
	data := SysData{
		CPUCores: sys.CPUCores(),
		CPUUsage: sys.CPUUsage(),
	}

	file, err := os.OpenFile("stats.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
}