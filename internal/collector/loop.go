package collector

import (
	"context"
	"prefect/internal/collect"
	"prefect/internal/state"
	"time"
)

func Run(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// prime CPU sampler (important)
	collect.CPUInfo()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			cpu := collect.CPUInfo()
			ramPct, ramUsed, ramTotal := collect.RAMInfo()

			state.UpdateCPU(cpu)
			state.UpdateRAM(ramPct, ramUsed, ramTotal)
		}
	}
}
