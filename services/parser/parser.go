package parser

import (
	"prefect/services/sys"
)

type SysData struct {
	// CPU
	CPUCores   int `json:"cpu_cores"`
	CPUThreads int `json:"cpu_threads"`
	CPUUsage   int `json:"cpu_usage"`
	CPUTemp    int `json:"cpu_temp"`

	// RAM
	RAMTotal int `json:"ram_total"`
	RAMUsed  int `json:"ram_used"`
	RAMUsage int `json:"ram_usage"`

	// Disk
	DiskTotal uint64 `json:"disk_total"`
	DiskUsed  uint64 `json:"disk_used"`
	DiskUsage int    `json:"disk_usage"`
}

func SysDataParser() SysData {
	RAMTotal, RAMUsed, RAMUsage := sys.RAM()
	DiskTotal, DiskUsed, DiskUsage := sys.Disk()

	// Data Structures
	return SysData{
		CPUCores:   sys.CPUCores(),
		CPUThreads: sys.CPUThreads(),
		CPUUsage:   sys.CPUUsage(),
		CPUTemp:    sys.CPUTemp(),
		RAMTotal:   RAMTotal,
		RAMUsed:    RAMUsed,
		RAMUsage:   RAMUsage,
		DiskUsed:   DiskUsed,
		DiskTotal:  DiskTotal,
		DiskUsage:  DiskUsage,
	}

}
