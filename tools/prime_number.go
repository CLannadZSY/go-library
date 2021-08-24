/*
素数, 没看懂的一节...
https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < math.MaxUint8; i++ {
			ch <- i
		}
	}()
	return ch
}
