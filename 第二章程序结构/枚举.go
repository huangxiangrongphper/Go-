package main

// 使用常量和iota关键字实现枚举。

import (
	"fmt"
)

const (
	a       = iota               // a == 0
	b                            // b == 1，隐式使用iota关键字，实际等同于 b = iota
	c                            // c == 2, 实际等同于 c = iota
	d, e, f = iota, iota, iota   // d=3, e=3, f=3, 同一行值相同，此处不能只写一个iota
	g       = iota
	h       = "h"
	i
	j       = iota
)














