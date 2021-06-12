package main

import "fmt"

func fibonacci() func() int {
	pre, cur := 0,1;
	return func() int {
		pre, cur = cur, pre+cur;
		return cur- pre; // Otherwise the series will start with 1; 
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
