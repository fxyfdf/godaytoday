package main

import (
	"fmt"
	"github.com/godaytoday/gopsutil_test/getmetrics"
)

func main() {
	fmt.Println("test")
	//getmetrics.GetCpuInfo()
	getmetrics.GetCpuLoad()
	//for i:=0; i < 100; i++{
	//	getmetrics.GetCpuLoad()
	//	//time.Sleep(time.Second * 2)
	//}
	getmetrics.GetMemInfo()
	getmetrics.GetHostInfo()
	getmetrics.GetDiskInfo()
	ip, err := getmetrics.GetLocalIP()
	if err != nil{
		fmt.Printf("local ip:%v",ip)
	}
	/**
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
**/
}

