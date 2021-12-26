package main

import (
	"fmt"
	"time"
)

func production_num(num_chan chan int) {
	for i := 1; i < 1000000; i++ {
		//Add 10000 numbers
		num_chan <- i
	}
	close(num_chan)

}

func cala_prime(num_chan, primeNum_chan chan int, exit_chan chan bool) {
	for {
		num, ok := <-num_chan
		flag := true
		if !ok {
			break
		}
		for j := 2; j < num; j++ {
			if num%j == 0 {
				//Not prime number
				flag = false
				break
			}

		}
		if flag {
			primeNum_chan <- num
		}

	}
	exit_chan <- true
}

func main() {
	//fmt.Println(time.Now().UnixNano())
	cache_time := time.Now()
	num_chan := make(chan int, 10000)
	primeNum_chan := make(chan int, 2000)
	exit_chan := make(chan bool, 10)

	go production_num(num_chan)

	for i := 1; i <= 10; i++ {
		go cala_prime(num_chan, primeNum_chan, exit_chan)
	}
	go func() {
		for i := 1; i <= 10; i++ {
			<-exit_chan
			//fmt.Println("Done", i)
		}
		close(primeNum_chan)
	}()

	for {
		//fmt.Println(" primeNum_chan length is", len(primeNum_chan))
		prime, ok := <-primeNum_chan
		if !ok {
			break
		}
		fmt.Println(prime)
	}
	now := time.Now()
	fmt.Println(now.Sub(cache_time))

} //参考链接https://blog.csdn.net/qq_42183962/article/details/103917626