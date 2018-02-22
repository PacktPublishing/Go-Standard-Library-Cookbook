package main

import (
	"fmt"
	"sync"
)

var names = []string{"Alan", "Joe", "Jack", "Ben",
	"Ellen", "Lisa", "Carl", "Steve", "Anton", "Yo"}

func main() {

	m := sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(idx int) {
			m.Store(fmt.Sprintf("%d", idx), names[idx])
			wg.Done()
		}(i)
	}
	wg.Wait()

	v, ok := m.Load("1")
	if ok {
		fmt.Printf("For Load key: 1 got %v\n", v)
	}

	v, ok = m.LoadOrStore("11", "Tim")
	if !ok {
		fmt.Printf("Key 11 missing stored val: %v\n", v)
	}

	m.Range(func(k, v interface{}) bool {
		key, _ := k.(string)
		t, _ := v.(string)
		fmt.Printf("For index %v got %v\n", key, t)
		return true
	})

}
