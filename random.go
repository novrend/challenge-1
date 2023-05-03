package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex

	var interface1 interface{} = [3]string{"bisa1", "bisa2", "bisa3"}
	var interface2 interface{} = [3]string{"coba1", "coba2", "coba3"}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go printArray(interface1, i, &wg, &mtx)
		wg.Add(1)
		go printArray(interface2, i, &wg, &mtx)
	}
	wg.Wait()
}

func printArray(array interface{}, i int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	fmt.Println(array, i)
	wg.Done()
}
