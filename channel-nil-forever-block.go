package main

import (
	"time"
	"fmt"
)


/**
WaitMany() looks like a good way to wait for channels a and b to close, but it has a problem.
Let’s say that channel a is closed first, then it will always be ready to receive.
Because bclosed is still false the program can enter an infinite loop,
preventing the channel b from ever being closed.
A safe way to solve the problem is to leverage the blocking properties of a nil channel and rewrite the program like this
In the rewritten WaitMany() we nil the reference to a or b once they have received a value.
When a nil channel is part of a select statement, it is effectively ignored,
so niling a removes it from selection, leaving only b which blocks until it is closed,
exiting the loop without spinning.

 */

//func WaitMany(a, b chan bool) {
//	var aclosed, bclosed bool
//	for !aclosed || !bclosed {
//		select {
//		case <-a:
//			aclosed = true
//		case <-b:
//			bclosed = true
//		}
//	}
//}

func WaitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
	}
}

func main() {
	a, b := make(chan bool), make(chan bool)
	t0 := time.Now()
	go func() {
		close(a)
		close(b)
	}()
	WaitMany(a, b)
	fmt.Printf("waited %v for WaitMany\n", time.Since(t0))
}
