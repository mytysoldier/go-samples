package main

import (
	"fmt"
	"sync"
	"time"
)

// func heavyProdess(num int) string {
// 	time.Sleep(1 * time.Second)
// 	return fmt.Sprintf("finish %d", num)
// }

// func main() {
// 	start := time.Now()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(heavyProdess(i + 1))
// 	}
// 	end := time.Now()

// 	fmt.Println("処理時間：", (end.Sub(start).Seconds()))
// }

func heavyProdess(num int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("finish %d", num)
}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go heavyProdess(i+1, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}
	end := time.Now()

	fmt.Println("処理時間：", (end.Sub(start).Seconds()))
}
