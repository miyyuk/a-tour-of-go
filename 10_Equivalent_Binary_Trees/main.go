package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		n1 := <-ch1
		n2 := <-ch2
		if n1 != n2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("tree.New(1)")
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		n := <-ch
		fmt.Print(n, " ")
	}

	fmt.Println("\ntree.New(2)")
	ch2 := make(chan int)
	go Walk(tree.New(2), ch2)
	for i := 0; i < 10; i++ {
		n := <-ch2
		fmt.Print(n, " ")
	}

	fmt.Println("\nSame")
	fmt.Print(Same(tree.New(1), tree.New(1)), " ")
	fmt.Print(Same(tree.New(1), tree.New(2)))
}

// >result
// tree.New(1)
// 1 2 3 4 5 6 7 8 9 10
// tree.New(2)
// 2 4 6 8 10 12 14 16 18 20
// Same
// true false
