package main

import (
	"fmt"
)

// 当变量被声明但没有赋值时，它的值就是定义时指定类型的零值。
// 例如int为0、float为0.0、bool为false、string为空字符串、指针为nil等。
// 如果变量甚至没有指定类型，那么系统默认它为int型，值为0，所以变量在Go语言中都是需要初始化的。

// 使用驼峰法命名变量：首个单词小写，每个独立单词的首字母大写，例如：setNum和backTop
// 常量命名建议全部为大写字母（遵循Shell环境的习惯）

// 命名规范并不是强制性规定，实际上在Go语言中，程序实体（常量、变量、函数）首字母大小写是有特殊含义的（外部可用性）。

var x int

var xx bool

var xxx string

var xxxx, xxxxx int // 定义两个变量xxxx、xxxxx为int型

// 多变量可以在同一行进行声明和赋值
var kk, kkk, kkkk int = 1, 2, 3

// 上面的是静态声明方式，除此之外还有动态声明方式，与常量一样，例如：
var rr, rrr, rrrr = "string", 1, true

// 与常量一样写成这种形式
var (
	xxxxxx, xxxxxxx int
	xxxxxxxx bool
	xxxxxxxxx string
)

func main()  {
	kk, kkk = kkk, kk // 交换两个变量的值
	x, xx = 5, true // 多变量可以在同一行进行赋值（只能在函数体内）
	fmt.Print(x, xx, xxx, xxxx, xxxxx, xxxxxx, xxxxxxx, xxxxxxxx, xxxxxxxxx, kk, kkk, kkkk, rr, rrr)
}

