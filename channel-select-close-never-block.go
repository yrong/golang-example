/**
A closed channel never blocks
A nice solution to this problem is to leverage the property that a closed channel is always ready to receive.
Using this property we can rewrite the program, now including 100 goroutines,
without having to keep track of the number of goroutines spawned, or correctly size the finish channel
As the behaviour of the close(finish) relies on signalling the close of the channel, not the value sent or received,
declaring finish to be of type chan struct{} says that the channel contains no value;
we’re only interested in its closed property.
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//finish := make(chan bool)
	//var done sync.WaitGroup
	//done.Add(1)
	//go func() {
	//	select {
	//	case <-time.After(1 * time.Hour):
	//	case <-finish:
	//	}
	//	done.Done()
	//}()
	//t0 := time.Now()
	//finish <- true // send the close signal
	//done.Wait()    // wait for the goroutine to stop
	//fmt.Printf("Waited %v for goroutine to stop\n", time.Since(t0))
	const n = 100
	finish := make(chan struct{})
	var done sync.WaitGroup
	for i := 0; i < n; i++ {
		done.Add(1)
		go func() {
			select {
			case <-time.After(1 * time.Hour):
			case <-finish:
			}
			done.Done()
		}()
	}
	t0 := time.Now()
	close(finish)    // closing finish makes it ready to receive
	done.Wait()      // wait for all goroutines to stop
	fmt.Printf("Waited %v for %d goroutines to stop\n", time.Since(t0), n)
}