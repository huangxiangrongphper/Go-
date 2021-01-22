package main

// 使用常量和iota关键字实现枚举。

import (
	"fmt"
)

const (
	a       = iota               // a == 0 iota是一个从0开始递增的整型值
	b                            // b == 1，隐式使用iota关键字，实际等同于 b = iota
	c                            // c == 2, 实际等同于 c = iota
	d, e, f = iota, iota, iota   // d=3, e=3, f=3, 同一行值相同，此处不能只写一个iota
	g       = iota               // g == 4
	h       = "h"                // h == "h"，单独赋值，iota依旧递增为5
	i                            // i == "h"，默认使用上面的赋值，iota依旧递增为6
	j       = iota               // j == 7
)

const z = iota // 每个单独定义的const常量中，iota都会重置，此时 z == 0

func main() {
	fmt.Print(a, b, c, d, e, f, g, h, i, j, z)
}

// 每个const定义的第一个常量被默认设置为0，显示设置为其他值除外。
// 后续的常量默认设置为它上面那个常量的值，如果前面那个常量的值是iota，
// 则它也被设置为iota。因为iota有递增效果，所以可以实现枚举。
// 此外，iota也可以用在表达式中，如：iota + 50.














