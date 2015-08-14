// 在一维坐标轴上有n个区间段，求重合区间最长的两个区间段

// 解决思路，先对坐标左端进行排序，由排序好的区间进行右端比较；
// 对x2 >= x1, 比y2与y1：
//		init: maxDist = 0
//      若 x2 >= y1，delete (x1,y1)
//      否则，若 y2 > y1, d|max = max(maxDist, y1-x2)，delete (x1,y1)
//		     若 y2 <= y1, d|max = max(maxDist, y2-x2), delete (x2,y2)
// return maxDist;

package main

import (
	"sort"
	"log"
	"math/rand"
)

func MaxInt(x, y int) int {
	if x >=y {
		return x
	}
	return y
}

type Coordin struct {
	X int
	Y int
}

type Coordins []Coordin

func (cords Coordins)Len() int {
	return len(cords)
}

func (cords Coordins)Less(i, j int) bool {
	return cords[i].X < cords[j].X
}

func (cords Coordins)Swap(i, j int) {
	cords[i], cords[j] = cords[j], cords[i]
}

func SortCoords(src []Coordin) {
	sort.Sort(Coordins(src))
}

// require cmp sort by X
func CompareAll(cmp []Coordin) (maxDist int) {

	if len(cmp) <= 1 {
		return maxDist
	}

	fst, sec := cmp[0], cmp[1]
	maxDist, vic := comparePair(fst, sec)

	log.Println("compare: ", cmp[0], " && ", cmp[1], " and current maxDist is ", maxDist)

	cmp1 := append([]Coordin{}, vic)
	cmp1 = append(cmp1, cmp[2:]...)

	return MaxInt(maxDist, CompareAll(cmp1))
}

func comparePair(fst, sec Coordin) (maxDist int, vic Coordin) {
	// x2 >= y1; 0, sec
	if sec.X >= fst.Y {
		return maxDist, sec
	}

	// x2 < y1 && y2 <= y1; y2-x2, fst
	if sec.Y <= fst.Y {
		return sec.Y-sec.X, fst
	}

	// x2 < y1 && y2 > y1; y1-x2, sec
	return fst.Y-sec.X, sec
}

// 初始化一组 []Coordin
func initCoordin(n int) (s []Coordin) {

	for i:=0; i<n; i++ {
		a := rand.Intn(100)
		s = append(s, Coordin{a, a + rand.Intn(100)})
	}

	return s
}

func main() {
	n := 10
	arr := initCoordin(n)

	log.Println("整理之前: ", arr)

	SortCoords(arr)

	log.Println("整理之后: ", arr)

	l := CompareAll(arr)

	log.Println("最长区段长为:", l)
}