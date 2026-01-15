package apps

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/cpu" // gopsutil package for CPU usage
	"time"
)

func CPUInfo() float64{
	for {
		cpuUsage, err := cpu.Percent(time.Second, false)

		if err != nil {
			fmt.Println("Error reading CPU:", err)
			return 0
		} // catch CPU error

        return cpuUsage[0]
        /*
		fmt.Printf("\rCPU Usage: %.2f%% | Cores: %d ", cpuUsage[0], cpuCores)
        */
	}
}

func CPUCores() int {
    cores, _ := cpu.Counts(true)
    return cores
}
