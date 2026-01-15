package apps

import (
	"github.com/shirou/gopsutil/v4/mem"
)

func RAMInfo() (float64, float64, float64) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, 0, 0
	}

	// Convert bytes to MiB (1024^2)
	usedMiB := float64(v.Used) / (1024 * 1024)
	totalMiB := float64(v.Total) / (1024 * 1024)

	return v.UsedPercent, usedMiB, totalMiB
}