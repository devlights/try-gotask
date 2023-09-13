package main

import "fmt"

func main() {
	var (
		ch = make(chan string)
	)

	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%d:helloworld", i)
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
