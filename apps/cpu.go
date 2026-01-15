package apps

import (
    "fmt"
    "github.com/shirou/gopsutil/v4/cpu"
    "time"
)

func CPUInfo() float64 {
    // 1. Just do the action ONCE
    cpuUsage, err := cpu.Percent(time.Second, false)

    if err != nil {
        fmt.Println("Error reading CPU:", err)
        return 0
    }

    // 2. Return the result and get out
    return cpuUsage[0]
}

func CPUCores() int {
    cores, _ := cpu.Counts(true)
    return cores
}