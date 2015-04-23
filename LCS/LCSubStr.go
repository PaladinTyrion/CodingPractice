// 最长公共子串。问题比最长公共子序列简单。
// 算法参见：
// http://en.wikibooks.org/wiki/Algorithm_Implementation/Strings/Longest_common_substring

package main

import (
	"log"
)

func LCSubStr(s1 string, s2 string) string {
	l1, l2 := len(s1), len(s2)

	var matrix = make([][]int, l1+1)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, l2+1)
	}

	longest := 0
	x_longest := 0
	for x := 1; x < l1+1; x++ {
		for y := 1; y < l2+1; y++ {
			if s1[x-1] == s2[y-1] {
				matrix[x][y] = matrix[x-1][y-1] + 1
				if matrix[x][y] > longest {
					longest = matrix[x][y]
					x_longest = x
				}
			} else {
				matrix[x][y] = 0
			}
		}
	}

	return s1[x_longest-longest : x_longest]
}

func main() {
	str1, str2 := "abasdfasdf", "fasdfasdfas"
	lcss := LCSubStr(str1, str2)
	log.Println(lcss)
}
