package parser

import (
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

func SysDataParser() SysData {
	RAMTotal, RAMUsed, RAMPercent := sys.RAM()

	// Data Structures
	return SysData{
		CPUCores:   sys.CPUCores(),
		CPUUsage:   sys.CPUUsage(),
		RAMTotal:   RAMTotal,
		RAMUsed:    RAMUsed,
		RAMPercent: RAMPercent,
	}

}
