package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Resource struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

type ResourceLabel struct {
	Label string  `json:"label"`
	Usage float64 `json:"usage"`
}

type ResourceInfoResponse struct {
	Memory   Resource        `json:"memory"`
	CpuUsage float64         `json:"cpuUsage"`
	CpuCores []ResourceLabel `json:"cpuCores"`
}

func getCPUUsagePercentage() float64 {
	percentages, _ := cpu.Percent(time.Second, false)

	return percentages[0]
}

func getCpuPercentages() []ResourceLabel {
	percentages, _ := cpu.Percent(time.Second, true)

	labels := make([]ResourceLabel, len(percentages))

	for index, percentage := range percentages {
		labels[index] = ResourceLabel{
			Label: fmt.Sprintf("cpu%v", index),
			Usage: percentage,
		}
	}

	return labels
}

func getMemInfo() Resource {
	mem, _ := mem.VirtualMemory()

	return Resource{
		Total: mem.Total / 1024 / 1024,
		Used:  mem.Used / 1024 / 1024,
	}
}

func getResourceInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	var wg sync.WaitGroup
	wg.Add(3)

	var memory Resource
	var cpuUsage float64
	var cpuCores []ResourceLabel

	go func() {
		defer wg.Done()
		memory = getMemInfo()
	}()

	go func() {
		defer wg.Done()
		cpuUsage = getCPUUsagePercentage()
	}()

	go func() {
		defer wg.Done()
		cpuCores = getCpuPercentages()
	}()

	wg.Wait()

	response := ResourceInfoResponse{
		Memory:   memory,
		CpuUsage: cpuUsage,
		CpuCores: cpuCores,
	}

	processes, _ := process.Processes()
	fmt.Println(len(processes))
	resJson, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}
