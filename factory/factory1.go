// 2个工厂分布在一条东西向高速公路的两侧，工厂距离公路最西端的距离
// 分别是0、4、5、10、12、18、27、30、31、38、39、47.
// 在这12个工厂中选取3个原料供应厂，使得剩余工厂到最近的原料
// 供应厂距离之和最短，问应该选哪三个厂 ？

// 暴力解决的方案
// output：min:43; Num:5,27,39
// 时间复杂度：O(n^4)

package main

import (
	"fmt"
)

//三数互不相等
func NotEqualEveryone(x, y, z int) bool {
	if x != y && x != z && y != z {
		return true
	}
	return false
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//两数最小
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

//三数最小
func Min3(x, y, z int) int {
	return Min2(Min2(x, y), z)
}

//离三点最近的值
func Nearest(x, y, z int, l int) int {
	return Min3(Abs(l-x), Abs(l-y), Abs(l-z))
}

//输入数组，输出最小距离，以及对应数组的三个数字
func GetValandOrd(src []int) (keepMin int, res [3]int) {
	s := len(src)
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			for k := 0; k < s; k++ {
				if NotEqualEveryone(i, j, k) {
					det := 0
					for l := 0; l < s; l++ {
						det += Nearest(src[i], src[j], src[k], src[l])
					}
					if keepMin == 0 {
						keepMin = det
					}
					if det < keepMin {
						keepMin = det
						res = [3]int{src[i], src[j], src[k]}
					}
				}
			}
		}
	}
	return keepMin, res
}

func main() {
	src := []int{0, 4, 5, 10, 12, 18, 27, 30, 31, 38, 39, 47}
	src1 := []int{0, 15, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47}

	keepMin, order := GetValandOrd(src)
	fmt.Printf("min:%v; Num:%v,%v,%v\n", keepMin, order[0], order[1], order[2])
	keepMin, order = GetValandOrd(src1)
	fmt.Printf("min:%v; Num:%v,%v,%v\n", keepMin, order[0], order[1], order[2])

}
