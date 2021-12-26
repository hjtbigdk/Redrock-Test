package main

import (
"fmt"
)

func main() {
	var x string
	ch := make(chan string,1)
	ch <- "下山的路又堵起了"
	x = <- ch

	go func() {
		fmt.Println(x)
	}()
}

