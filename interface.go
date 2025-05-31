package golist

type SortedList interface {
	Distinct() SortedList
	At(int) int
	Insert(v int)
	Len() int
}
