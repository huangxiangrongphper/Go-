// 定义当前代码所属的包，作为模块化的标识
// main包是一个特殊的包名，表示当前是一个可执行程序，而不是一个库
package main

// 要导入的代码包
// Go语言要求只有用到的包才能导入，如果导入一个代码包又不使用，那么编译时会出错。
import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
}
