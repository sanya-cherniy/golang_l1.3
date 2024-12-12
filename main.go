package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [...]int{2, 4, 6, 8, 10}
	var res int
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := range arr {
		wg.Add(1)
		go func(x int, res *int) {
			x_sqr := x * x
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			*res += x_sqr
		}(arr[i], &res)
	}
	wg.Wait()
	fmt.Println(res)
}
