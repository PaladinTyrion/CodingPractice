package main

import (
	"log"
)

// 获取n1-n2之间的所有偶数
func even(n1, n2 int) (array []int) {
	for i := n1; i <= n2; i++ {
		if i&1 == 0 { // 位操作符&与C语言中使用方式一样
			array = append(array, i)
		}
	}
	return
}

// 互换两个变量的值
// 不需要使用第三个变量做中间变量
// 原理：异或本质上记录两个二进制不相同的位信息。只要有这份记录，外加随便两数之一，都可还原另外一数
// 注意这里空间复杂度为O(1),在考虑需要空间复杂度为O(1),可以看看是否可使用位运算达到目的
func swap(a, b int) (int, int) {
	a ^= b // 异或等于运算
	b ^= a
	a ^= b
	return a, b
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

// 变换符号
func nagation(a int) int {
	// 注意: C语言中是 ~a+1这种方式
	return ^a + 1 // Go语言取反方式和C语言不同，Go语言不支持~符号。 ^0 = -1
}

func main() {
	a, b := swap(2, 5)
	log.Printf("swap: %d\t%d\n", a, b)
}
