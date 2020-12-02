package main

import "github.com/jlevitt/katas/bst"

func main() {
	b := bst.Bst{}
	b.Insert(3, "a")
	b.Insert(5, "b")
	b.Insert(7, "c")
	b.Insert(2, "d")
	b.Insert(1, "e")
	b.Insert(4, "f")
	b.Insert(10, "g")
}
