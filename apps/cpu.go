package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu" // gopsutil package for CPU usage
	"time"
)

func main() {
	cpuCores, err := cpu.Counts(true)

	if err != nil {
		fmt.Println("Error getting CPU core count:", err)
		return
	} // catch CPU error

	for {
		cpuUsage, err := cpu.Percent(time.Second, false)

		if err != nil {
			fmt.Println("Error reading CPU:", err)
			return
		} // catch CPU error

		fmt.Printf("\rCPU Usage: %.2f%% | Cores: %d ", cpuUsage[0], cpuCores)
	}
}
