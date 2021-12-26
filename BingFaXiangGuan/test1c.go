package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		ch <- 1
	}
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	go func() {
		<-ch
		fmt.Println(123)
	}()
	wg.Wait()
}
