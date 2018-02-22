package main

import "sync"
import "fmt"

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			// Do some work
			defer wg.Done()
			fmt.Printf("Exiting %d\n", idx)
		}(i)
	}
	wg.Wait()
	fmt.Println("All done.")
}
