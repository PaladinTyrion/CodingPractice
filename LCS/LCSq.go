// 最长公共子序列，非最长公共子串。
// 算法参见 http://www.julyedu.com/video/play/id/9
// 需要注意的是：当存在多个公共子序列时，优先顺从str1的顺序，
// 该顺序由lcsMatrix[i-1][j].Count >= lcsMatrix[i][j-1].Count决定，
// 另外，x,y记录继承上一次公共字符的位置，由于str1[x]==str2[y]，故可只记录一个

package main

import (
	"log"
)

type Inherit struct {
	x int
	//	y int
}

type Record struct {
	Inherit
	Count int
}

func LCSq(str1, str2 string) (lenth int, lcs string) {
	l1, l2 := len(str1), len(str2)

	lcsMatrix := make([][]Record, l1+1)
	for i := 0; i < len(lcsMatrix); i++ {
		lcsMatrix[i] = make([]Record, 1+l2)
	}

	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			tmpLM := &lcsMatrix[i][j]
			if str1[i-1] == str2[j-1] {
				tmpLM.Count = lcsMatrix[i-1][j-1].Count + 1
				tmpLM.x = i - 1
				//				tmpLM.y = j - 1
				continue
			}

			if lcsMatrix[i-1][j].Count >= lcsMatrix[i][j-1].Count {
				*tmpLM = lcsMatrix[i-1][j]
			} else {
				*tmpLM = lcsMatrix[i][j-1]
			}
		}
	}

	lenth = lcsMatrix[l1][l2].Count
	bLCS := make([]byte, lenth)
	for k := lenth - 1; k >= 0; k-- {
		tmp := lcsMatrix[l1][l2]
		bLCS[k] = str1[tmp.x]
		l1 = tmp.x
		//		l2 = tmp.y
	}
	lcs = string(bLCS)

	return lenth, lcs
}

func main() {

	str1 := "BADCDCBA"
	str2 := "ABCDCDAB"

	lenth, lcs := LCSq(str1, str2)
	log.Println(lenth, lcs)
}
