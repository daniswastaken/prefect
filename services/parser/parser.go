package parser

import (
	"encoding/json"
	"log"
	"os"
	"prefect/services/sys"
)

type SysData struct {
	// CPU
	CPUCores int `json:"cpu_cores"`
	CPUUsage int `json:"cpu_usage"`

	// RAM
	RAMTotal   int `json:"ram_total"`
	RAMUsed    int `json:"ram_used"`
	RAMPercent int `json:"ram_percent"`
}

func SysDataParser() {
	RAMTotal, RAMUsed, RAMPercent := sys.RAM()

	// Data Structures
	data := SysData{
		CPUCores:   sys.CPUCores(),
		CPUUsage:   sys.CPUUsage(),
		RAMTotal:   RAMTotal,
		RAMUsed:    RAMUsed,
		RAMPercent: RAMPercent,
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
