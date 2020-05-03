package getmetrics

import (
	"fmt"
	"github.com/shirou/gopsutil/load"
)

func GetCpuLoad(){
	info,_ := load.Avg()
	fmt.Printf("%v\n",info)

}
