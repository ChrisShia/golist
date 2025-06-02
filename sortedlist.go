package golist

import (
	"github.com/ChrisShia/gosort"
)

type sortedList []int

func (sl *sortedList) Len() int {
	return len(*sl)
}

func (sl *sortedList) Cap() int {
	return cap(*sl)
}

func (sl *sortedList) Insert(v int) {
	insertPosition := bisect(*sl, v)
	insert(sl, v, insertPosition)
}

func insert(oi *sortedList, v int, insertPosition int) {
	res := make(sortedList, len(*oi)+1)
	copy(res, (*oi)[:insertPosition])
	res[insertPosition] = v
	copy(res[insertPosition+1:], (*oi)[insertPosition:])
	*oi = res
}

func prepend(oi sortedList, v int) sortedList {
	tmp := make([]int, len(oi)+1)
	copy(tmp[1:], oi)
	tmp[0] = v
	return tmp
}

func MakeSortedList(size ...int) SortedList {
	var list []int
	switch len(size) {
	case 0:
		list = make([]int, 0)
	case 1:
		list = make([]int, size[0])
	default:
		list = make([]int, size[0], size[1])
	}
	sl := sortedList(list)
	return &sl
}

func MakeSorted(list []int) SortedList {
	sl := sortedList(list)
	if len(list) > 0 {
		sl.Sort()
	}
	return &sl
}

const ZEROTH int = 0

func bisect(oi sortedList, v int) int {
	length := len(oi)
	if length == 0 {
		return ZEROTH
	}
	if v >= oi[length-1] {
		return length
	}
	if v <= oi[ZEROTH] {
		return ZEROTH
	}

	left, right := leftRightHalves(oi)

	var position int
	if v > right[ZEROTH] {
		position = len(left) + bisect(right, v)
	} else if v == right[ZEROTH] {
		position += len(left)
	} else if len(left) == 0 {
		return position
	} else if v == left[len(left)-1] {
		position = len(left)
	} else if v > left[ZEROTH] {
		position = bisect(left, v)
	} else {
		position = ZEROTH
	}
	return position
}

func leftRightHalves(oi sortedList) (sortedList, sortedList) {
	halfLength := len(oi) / 2
	left := oi[:halfLength]
	right := oi[halfLength:]
	return left, right
}

func (sl *sortedList) Sort() {
	gosort.MergeSortInPlace(*sl)
}

func join(left, right sortedList) sortedList {
	res := make(sortedList, len(left)+len(right))
	copy(res, left)
	copy(res[len(left):], right)
	return res
}

func (sl *sortedList) Distinct() SortedList {
	return sl.distinct()
}

func (sl *sortedList) distinct() *sortedList {
	uniques := make(sortedList, 0)
	for _, e := range *sl {
		if len(uniques) == 0 || e != uniques[len(uniques)-1] {
			uniques = append(uniques, e)
		}
	}
	return &uniques
}

func (sl *sortedList) At(i int) int {
	return (*sl)[i]
}
