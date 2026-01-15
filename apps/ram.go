package apps

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/mem"
)

func RAMInfo() float64 {
	// This one call gets EVERYTHING (Total, Used, Free, Percent, etc.)
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error getting RAM info:", err)
		return 0
	}

	return v.UsedPercent
	
	/*
	// Now you access the fields from the variable 'v'
	fmt.Printf("RAM Usage: %.2f%% | Total RAM: %v Bytes\n", v.UsedPercent, v.Total)
	
	// Better yet, let's make the RAM size readable in GB
	gb := v.Total / 1024 / 1024 / 1024
	fmt.Printf("Simplified: %.2f%% of %d GB used\n", v.UsedPercent, gb)
	*/
}