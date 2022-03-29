package main

import (
	"log"
	"time"

	"yandex-practicum-go-devops/internal/collector"
	"yandex-practicum-go-devops/internal/sender"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

func main() {
	collector := collector.NewCollector()
	sender := sender.NewSender("http://127.0.0.1:8080/update")

	//start tikers
	pollTicker := time.NewTicker(pollInterval)
	reportTiker := time.NewTicker(reportInterval)

	for {
		select {
		case <-pollTicker.C:
			log.Println("Poll")
			collector.Collect()
		case <-reportTiker.C:
			log.Println("Report")
			sender.SendGauge(collector.Gauges)
			sender.SendCounter(collector.Counters)
		}
	}
}
