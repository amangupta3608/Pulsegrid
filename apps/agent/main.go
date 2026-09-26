package main
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Metric struct {
	Hostname string
	Timestamp time.Time
	CPUPercent float64
	MemUsedMB uint64
	MemTotalMB uint64
}

func collectMetrics() Metric{
	percentage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Fatal(err)
	}

	memStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	return Metric{
		Hostname: hostname,
		Timestamp: time.Now(),
		CPUPercent: percentage[0],
		MemUsedMB: memStat.Used / 1024 / 1024,
		MemTotalMB: memStat.Total / 1024 / 1024,
	}
}

func main() {
	metric := collectMetrics()
	fmt.Printf("%+v\n", metric) 
}