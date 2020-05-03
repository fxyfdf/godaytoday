package main //声明包名，如果包名是main 包 说明包是生成可执行文件
// 如果不是main 包，说明包是给 其他包用的，属于工具包类
//声明这个文件属于哪个包

// 导入包语句
import (
	"fmt"
	"github.com/godaytoday/helloworld/hello"
)

// 函数外只能放置标识符 (变量，常量，函数，类型)的声明
// fmt.Println("非法，无法执行")

// 程序入口函数

func main() { // 整个函数的入口 无参数和返回值
	fmt.Println("hello world")
	// 调用包时，函数需要大写
	hello.HlloTest()

}
