// https://go-tour-jp.appspot.com/moretypes/26

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(i int) int {
	fibo := []int{0, 1}
	return func(i int) int {
		if i > 1 {
			fibo = append(fibo, fibo[i-2]+fibo[i-1])
		}
		return fibo[i]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
