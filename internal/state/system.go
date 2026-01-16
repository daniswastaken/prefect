package state

import (
	"sync"
	"time"
)

type SystemStats struct {
	CPUUsage float64
	Cores    int

	RAMPct   float64
	RAMUsed  float64
	RAMTotal float64

	UpdatedAt time.Time
}

var (
	mu    sync.RWMutex
	stats SystemStats
)

func Init(cores int) {
	mu.Lock()
	stats.Cores = cores
	stats.UpdatedAt = time.Now()
	mu.Unlock()
}

func UpdateCPU(usage float64) {
	mu.Lock()
	stats.CPUUsage = usage
	stats.UpdatedAt = time.Now()
	mu.Unlock()
}

func UpdateRAM(pct, used, total float64) {
	mu.Lock()
	stats.RAMPct = pct
	stats.RAMUsed = used
	stats.RAMTotal = total
	stats.UpdatedAt = time.Now()
	mu.Unlock()
}

func Snapshot() SystemStats {
	mu.RLock()
	defer mu.RUnlock()
	return stats
}
