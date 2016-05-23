package main

import "errors"

func SearchKeyByDichotomy(orderedArr []int, start int, end int, key int) (error, int) {
	if start > end {
		return errors.New("Your start should be smaller than end."), -1;
	}
	if end >= len(orderedArr) {
		return SearchKeyByDichotomy(orderedArr, start, len(orderedArr)-1, key);
	}

	mid := (start+end)/2;
	if orderedArr[mid] == key {
		return nil, mid;
	}

	if start == end {
		return nil, -1;
	}

	if key > orderedArr[mid] {
		return SearchKeyByDichotomy(orderedArr, mid+1, end, key);
	}

	return SearchKeyByDichotomy(orderedArr, start, mid-1, key);
}

func main() {

	arr := []int{1, 36, 38, 56, 99};

	err , pos := SearchKeyByDichotomy(arr, 0, len(arr), 99);
	if err != nil {
		println(err.Error());
	}

	println(pos);
}
