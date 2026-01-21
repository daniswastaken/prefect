package sys

import (
	"github.com/shirou/gopsutil/v4/mem"
)


func RAM() (int, int, int) {
	virtualMem, err := mem.VirtualMemory()

	if err != nil {
		return 0, 0, 0
	}

	total := int(virtualMem.Total / (1024 * 1024))
	used := int(virtualMem.Used / (1024 * 1024))
	percent := int(virtualMem.UsedPercent)

	return total, used, percent
}
