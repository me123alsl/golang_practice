package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {

	// 변수 설정
	var playTime int = 5 // Sleep Time (Second)
	var wgCount int = 5  // Wait Group Count (Go Routine Count)

	// Wait Group 생성
	wg := sync.WaitGroup{}
	wg.Add(wgCount)

	// Go Routine 생성
	for i := 1; i <= wgCount; i++ {
		go newFunction(playTime, strconv.Itoa(i), &wg)
	}

	// Defer
	// Wait Group 대기
	defer func() {
		wg.Wait()
		fmt.Println("FINISH")
	}()

}

func newFunction(sleepSecond int, msg string, wg *sync.WaitGroup) {
	fmt.Println(msg)
	time.Sleep(time.Duration(sleepSecond) * time.Second)
	defer wg.Done()
}
