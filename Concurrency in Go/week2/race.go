/*Write two goroutines which have a race condition when executed concurrently.
 *Explain what the race condition is and how it can occur.
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	/* A race condition is the behavior of an electronics, software,
	 * or other system where the system's substantive behavior is dependent on
	 * the sequence or timing of other uncontrollable events

	 * Races occur due to communication
	 */
	for i := 0; i < 5; i++ {
		var my_num int = 20
		go addTen(&my_num)
		go minusTen(&my_num)
		time.Sleep(500)
		fmt.Printf("%v from main\n", my_num)
	}
}

func addTen(num *int) {
	for i := 0; i < 10; i++ {
		*num++
	}
	fmt.Printf("%v from AddTen\n", *num)
}

func minusTen(num *int) {
	for i := 0; i < 10; i++ {
		*num--
	}
	fmt.Printf("%v from MinusTen\n", *num)
}
