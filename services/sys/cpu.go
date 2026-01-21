package sys

import (
	"github.com/shirou/gopsutil/v4/cpu"
	"time"
)

func CPUCores() int {
	cores, err := cpu.Counts(true)

	if err != nil {
		return 0
	}

	return cores
}

func CPUUsage() int {
	percentages, err := cpu.Percent(time.Second, false)

	if err != nil || len(percentages) == 0 {
		return 0
	}

	return int(percentages[0])
}
