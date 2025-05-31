package golist

import "testing"

func TestSortedList_Insert(t *testing.T) {
	var tt = []struct {
		name         string
		input        int
		orderedSlice []int
		expected     []int
	}{
		{"", 5, []int{}, []int{5}},
		{"", 0, []int{1}, []int{0, 1}},
		{"", 2, []int{1}, []int{1, 2}},
		{"", 5, []int{1, 3, 7}, []int{1, 3, 5, 7}},
		{"", 5, []int{1, 3, 7, 7, 7, 10}, []int{1, 3, 5, 7, 7, 7, 10}},
		{"", 7, []int{1, 3, 7, 7, 7, 10}, []int{1, 3, 7, 7, 7, 7, 10}},
		{"", 7, []int{7, 7, 7, 7, 7, 7, 7, 7}, []int{7, 7, 7, 7, 7, 7, 7, 7, 7}},
		{"", 7, []int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 7}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			oi := sortedList(tc.orderedSlice)
			oi.Insert(tc.input)
			for i := range oi {
				if oi[i] != tc.expected[i] {
					t.Errorf("got %v, expected %v", oi, tc.expected)
				}
			}
		})
	}
}

func TestSortedList_bisect(t *testing.T) {
	var tt = []struct {
		name         string
		input        int
		orderedSlice []int
		expected     int
	}{
		{"", 5, []int{}, 0},
		{"", 0, []int{1}, 0},
		{"", 2, []int{1}, 1},
		{"", 5, []int{1, 3, 7}, 2},
		{"", 5, []int{1, 3, 7, 7, 7, 10}, 2},
		{"", 7, []int{1, 3, 7, 7, 7, 10}, 3},
		{"", 7, []int{7, 7, 7, 7, 7, 7, 7, 7}, 8},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			oi := sortedList(tc.orderedSlice)
			index := bisect(oi, tc.input)
			if index != tc.expected {
				t.Errorf("got %v, expected %v", index, tc.expected)
			}
		})
	}
}

func TestSortedList_Order(t *testing.T) {
	var tt = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"", []int{1, 4, 3, 5, 2}, []int{1, 2, 3, 4, 5}},
		{"", []int{10, 1, 1, 9, 6, 1, 1}, []int{1, 1, 1, 1, 6, 9, 10}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			oi := sortedList(tc.input)
			oi.Sort()
			for i := range oi {
				if oi[i] != tc.expected[i] {
					t.Errorf("got %v, epxected %v", oi[i], tc.expected[i])
				}
			}
		})
	}
}

func Test_LeftRightHalves(t *testing.T) {
	var tt = []struct {
		name                        string
		input                       []int
		expectedLeft, expectedRight []int
	}{
		{"", []int{1}, []int{}, []int{1}},
		{"", []int{1, 4, 3, 5, 2}, []int{1, 4}, []int{3, 5, 2}},
		{"", []int{1, 4, 3, 3, 5, 2}, []int{1, 4, 3}, []int{3, 5, 2}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			oi := sortedList(tc.input)
			left, right := leftRightHalves(oi)
			for i := range left {
				if left[i] != tc.expectedLeft[i] {
					t.Errorf("for left half got %v, expected %v", left[i], tc.expectedLeft[i])
				}
			}
			for i := range right {
				if right[i] != tc.expectedRight[i] {
					t.Errorf("for right half got %v, expected %v", right[i], tc.expectedRight[i])
				}
			}
		})
	}
}

func Test_SortedListDistinct(t *testing.T) {
	var tt = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"", []int{1, 2, 2, 2, 2, 4}, []int{1, 2, 4}},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			oi := sortedList(tc.input)
			oi.Sort()
			distinctList := oi.distinct()
			for i := range *distinctList {
				if (*distinctList)[i] != tc.expected[i] {
					t.Errorf("got %v, expected %v", (*distinctList)[i], tc.expected[i])
				}
			}
		})
	}
}
