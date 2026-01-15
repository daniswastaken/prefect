package main

import (
	"fmt"
	"time"
	"prefect/apps"
)

func main() {
    cores := apps.CPUCores()

    for {
        usage := apps.CPUInfo() // Get this every second
        ram := apps.RAMInfo()

        fmt.Printf("\rCores: %d | CPU: %.2f%% | RAM: %.2f%%", cores, usage, ram)
        time.Sleep(time.Second)
    }
}