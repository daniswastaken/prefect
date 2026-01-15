package apps

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
)

func RAMInfo() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting RAM info:", err)
		return 0
	}

	return v.UsedPercent
}