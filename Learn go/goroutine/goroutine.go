package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Task", id, "is running")
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)

		// go func (i int)  {
		// 	fmt.Println(i)
		// }(i)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
}
