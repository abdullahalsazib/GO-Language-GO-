package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task ", id)
}

func main() {
	fmt.Println("welcome in go on gorouting-->")
	var wg = sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
		// func(i int) {
		//     fmt.Println(i)
		// }(i)
	}

	wg.Wait()
}
