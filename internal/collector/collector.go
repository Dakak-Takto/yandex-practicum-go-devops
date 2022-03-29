package collector

import (
	"math/rand"
	"runtime"
)

type Collector interface {
	Collect()
}

type collector struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

var _ Collector = &collector{}

type Result struct {
	Key, Value, Type string
}

func NewCollector() *collector {
	c := collector{
		Gauges:   make(map[string]float64),
		Counters: make(map[string]int64),
	}
	return &c
}

func (c *collector) Collect() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	c.Counters["RandomValue"] = rand.Int63()
	c.Counters["PollCount"]++

	c.Gauges["Alloc"] = float64(m.Alloc)
	c.Gauges["BuckHashSys"] = float64(m.BuckHashSys)
	c.Gauges["Frees"] = float64(m.Frees)
	c.Gauges["GCCPUFraction"] = float64(m.GCCPUFraction)
	c.Gauges["GCSys"] = float64(m.GCSys)
	c.Gauges["HeapAlloc"] = float64(m.HeapAlloc)
	c.Gauges["HeapIdle"] = float64(m.HeapIdle)
	c.Gauges["HeapInuse"] = float64(m.HeapInuse)
	c.Gauges["HeapObjects"] = float64(m.HeapObjects)
	c.Gauges["HeapReleased"] = float64(m.HeapReleased)
	c.Gauges["HeapSys"] = float64(m.HeapSys)
	c.Gauges["LastGC"] = float64(m.LastGC)
	c.Gauges["Lookups"] = float64(m.Lookups)
	c.Gauges["MCacheInuse"] = float64(m.MCacheInuse)
	c.Gauges["MCacheSys"] = float64(m.MCacheSys)
	c.Gauges["MSpanInuse"] = float64(m.MSpanInuse)
	c.Gauges["MSpanSys"] = float64(m.MSpanSys)
	c.Gauges["Mallocs"] = float64(m.Mallocs)
	c.Gauges["NextGC"] = float64(m.NextGC)
	c.Gauges["NumForcedGC"] = float64(m.NumForcedGC)
	c.Gauges["NumGC"] = float64(m.NumGC)
	c.Gauges["OtherSys"] = float64(m.OtherSys)
	c.Gauges["PauseTotalNs"] = float64(m.PauseTotalNs)
	c.Gauges["StackInuse"] = float64(m.StackInuse)
	c.Gauges["StackSys"] = float64(m.StackSys)
	c.Gauges["Sys"] = float64(m.Sys)
	c.Gauges["TotalAlloc"] = float64(m.TotalAlloc)
}
