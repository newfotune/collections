package main

import (
	"fmt"

	"github.com/newfotune/collections/collections"
)

var comparator = func(a interface{}, b interface{}) int {
	var left int = a.(int)
	var right int = b.(int)

	return left - right
}

func main() {
	tree := collections.NewBst(comparator)
	tree.Add(1)
	tree.Add(2)
	tree.Add(-3)
	//tree.Add(4)

	fmt.Printf("%s ", tree)
}
