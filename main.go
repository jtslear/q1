package main

import (
	"fmt"
)

func theQuestion(i int, l []int) bool {

	for a := 0; a <= len(l); a++ {
		b := l[a]
		for c := 0; c <= len(l); c++ {
			if b+c == i {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("vim-go")
}
