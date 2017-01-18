package mergesort

import "testing"

func TestSimple(t *testing.T) {
	input := []int{4, 7, 3, 9, 1, 2}
	correct := []int{1, 2, 3, 4, 7, 9}
	output := Sort(input)
	for i := range correct {
		if correct[i] != output[i] {
			t.Fatal("sorting failed... expected:", correct, "received:", output)
		}
	}
}

func TestRepeatInts(t *testing.T) {
	input := []int{4, 4, 9, 9, 2, 2}
	correct := []int{2, 2, 4, 4, 9, 9}
	output := Sort(input)
	for i := range correct {
		if correct[i] != output[i] {
			t.Fatal("sorting failed... expected:", correct, "received:", output)
		}
	}
}
