package main

import (
	"fmt"

	"github.com/fupslot/go-sorting/internal/find"
)

func main() {
	r := find.DuplicatesN([]int{3, 4, 1, 2, 2, 2, 4, 5, 5})

	fmt.Println(r)
}
