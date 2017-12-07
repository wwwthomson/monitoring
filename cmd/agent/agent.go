package main

import (
	"github.com/wwwthomson/monitoring/pkg/agent"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	cpu := agent.CPU{}
	go cpu.RunJob(&wg)

	memory := agent.Memory{}
	go memory.RunJob(&wg)

	//swap := agent.Swap{}
	//go swap.RunJob(&wg)

	//disk := agent.Disk{}
	//go disk.RunJob(&wg)

	network := agent.Network{}
	go network.RunJob(&wg)

	wg.Wait()
}
