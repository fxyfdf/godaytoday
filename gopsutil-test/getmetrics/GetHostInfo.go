package getmetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/host"
)

// host info

func GetHostInfo() {
	hInfo, _ :=host.Info()
	fmt.Printf("host info: %v  boottime:%v\n",hInfo,hInfo.Uptime,hInfo.BootTime)

}