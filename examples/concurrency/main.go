package main

import (
	"fmt"
	"time"
	// "sync"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Microsecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}

	// Channel open/close
	// c := make(chan string)
	// go count("sheep", c)

	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// for {
	// 	msg, open := <- c

	// 	if !open {
	// 		break
	// 	}

	// 	fmt.Println(msg)
	// }

	// WaitGroup
	//  var wg sync.WaitGroup
	//  wg.Add(1)
	
	//  go func() {
	// 	 count("sheep")
	// 	 wg.Done()
	// }()
		 
	// wg.Wait()
	
	// GoRoutine
	// go count("sheep")
	// go count("fish")

	// time.Sleep(time.Second * 2)
	// fmt.Scanln()
}

// func count(thing string, c chan string) {
// 	for i := 1; i <= 5; i++ {
// 		// fmt.Println(i, thing)
// 		c <- thing
// 		time.Sleep(time.Millisecond * 500)
// 	}

// 	close(c)
// }