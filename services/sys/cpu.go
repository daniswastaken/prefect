package sys

import (
	"github.com/shirou/gopsutil/v3/host" // temp sensors only ava on v3
	"github.com/shirou/gopsutil/v4/cpu"
	"strings"
	"time"
	"strconv"
	"os"
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

	percent := int(percentages[0])

	if percent < 1 {
		return 1
	} else {
		return percent
	}
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

func readEnergyUJ(path string) (int64, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(strings.TrimSpace(string(data)), 10, 64)
}

func CPUPower(duration time.Duration) int {
	path := "/sys/class/powercap/intel-rapl:0/energy_uj"

	e1, err := readEnergyUJ(path)
	if err != nil {
		return 0
	}

	time.Sleep(duration)

	e2, err := readEnergyUJ(path)
	if err != nil {
		return 0
	}

	deltaUJ := e2 - e1
	seconds := duration.Seconds()

	// µJ → J → W
	joules := float64(deltaUJ) / 1_000_000.0
	watts := joules / seconds

	return int(watts)
}