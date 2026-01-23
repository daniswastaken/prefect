package sys

import (
	"math"
	"github.com/shirou/gopsutil/v4/disk"
)

func Disk() (uint64, uint64, int) {
    partitions, err := disk.Partitions(false)
    if err != nil {
        return 0, 0, 0
    }

    var totalBytes uint64
    var usedBytes uint64
    seenDevices := make(map[string]bool)

    for _, p := range partitions {
        if seenDevices[p.Device] {
            continue
        }

        usage, err := disk.Usage(p.Mountpoint)
        if err != nil {
            continue
        }

        totalBytes += usage.Total
        usedBytes += usage.Used
        seenDevices[p.Device] = true
    }

    if totalBytes == 0 {
        return 0, 0, 0
    }

    // Convert capacities to GiB
    total := totalBytes / (1024 * 1024 * 1024)
    used := usedBytes / (1024 * 1024 * 1024)

    // Calculate the used percentage
    percent := (float64(usedBytes) / float64(totalBytes)) * 100

    return total, used, int(math.Ceil(percent))
}