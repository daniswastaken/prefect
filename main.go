package main

import (
	"fmt"
	"time"
	"prefect/apps"
)

func main() {
    cores := apps.CPUCores()

    for {
        usage := apps.CPUInfo()
        ramPercentage, ramUsed, ramTotal := apps.RAMInfo()

        fmt.Printf("\rCores: %d | CPU: %.2f%% | RAM: %.2f%% (Used: %.0f MiB / Total: %.0f MiB)", cores, usage, ramPercentage, ramUsed, ramTotal)
        time.Sleep(time.Second)
    }
}