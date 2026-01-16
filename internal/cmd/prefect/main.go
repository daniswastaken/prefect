package main

import (
	"context"
	"fmt"
	"prefect/internal/collect"
	"prefect/internal/collector"
	"prefect/internal/state"
	"prefect/internal/api"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cores := collect.CPUCores()
	state.Init(cores)

	go collector.Run(ctx, time.Second)

	fmt.Println("Prefect starting at http://localhost:8080")
	api.StartHTTP(":8080")

	select {}
}