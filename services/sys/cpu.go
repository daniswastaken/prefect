package sys

import (
	"github.com/shirou/gopsutil/v3/host" // temp sensors only ava on v3
	"github.com/shirou/gopsutil/v4/cpu"
	"strings"
	"time"
)

func CPUCores() int {
	cores, err := cpu.Counts(false)

	if err != nil {
		return 0
	}

	return cores
}

func CPUThreads() int {
	threads, err := cpu.Counts(true)

	if err != nil {
		return 0
	}

	return threads
}

func CPUUsage() int {
	percentages, err := cpu.Percent(time.Second, false)

	if err != nil || len(percentages) == 0 {
		return 0
	}

	return int(percentages[0])
}

func CPUTemp() int {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		return 0
	}

	for _, t := range temps {
		key := strings.ToLower(t.SensorKey)

		// Intel + AMD common package sensors
		if strings.Contains(key, "package") ||
			strings.Contains(key, "coretemp") ||
			strings.Contains(key, "k10temp") {

			if t.Temperature > 0 {
				return int(t.Temperature)
			}
		}
	}

	return 0
}
