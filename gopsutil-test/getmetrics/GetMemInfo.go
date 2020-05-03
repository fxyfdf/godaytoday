package getmetrics

import "github.com/shirou/gopsutil/mem"
import "fmt"
// mem info
func GetMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info:%v\n", memInfo)
}