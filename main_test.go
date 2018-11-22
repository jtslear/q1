package main

import (
	"testing"
)

type Test struct {
	k             int
	list          []int
	desiredResult bool
}

var td []Test

func TestTheQuestion(t *testing.T) {
	td = append(td, Test{
		17,
		[]int{10, 15, 3, 7},
		true,
	})

	td = append(td, Test{
		19,
		[]int{10, 15, 3, 7},
		false,
	})

	for _, v := range td {
		result := theQuestion(v.k, v.list)

		if !result {
			t.Logf("Expected %v, got %v\n", v.desiredResult, result)
			t.Fail()
		}
	}
}
