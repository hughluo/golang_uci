/* Implement the dining philosopher’s problem with the following constraints/modifications.

 * There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
 * Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
 * The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
 * In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 * The host allows no more than 2 philosophers to eat concurrently.
 * Each philosopher is numbered, 1 through 5.
 * When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
 * where <number> is the number of the philosopher.
 * When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, \
 * where <number> is the number of the philosopher.
 */
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wgH sync.WaitGroup
	var wgP sync.WaitGroup
	host := Host{}
	csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		csticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i, csticks[i], csticks[(i+1)%5]}
	}

	apply := make(chan int, 2)
	permits := make([]chan bool, 5)
	eatens := make([]chan bool, 5)
	host.eatTime = make(map[int]int, 0)

	for i := 0; i < 5; i++ {
		permits[i] = make(chan bool, 1)
		eatens[i] = make(chan bool, 1)
		host.eatTime[i] = 0
	}
	host.start(apply, permits, eatens, &wgH)

	/* 	go host.ExecApply(apply, permits, &wg)
	   	for i := 0; i < 5; i++ {
	   		go host.ExecEaten(i, eatens[i], &wg)
	   	} */
	for _, p := range philos {
		p.start(permits, eatens, apply, &wgP)
	}
	//time.Sleep(time.Second * 20)
	wgP.Wait()
}

type Host struct {
	// eatTime[2] == 3 means philo no.2 has eaten 3 times.
	eatTime map[int]int
}

func (h Host) start(apply chan int, permits []chan bool, eatens []chan bool, wg *sync.WaitGroup) {
	go h.execApply(apply, permits, wg)

	for i := 0; i < 5; i++ {
		go h.execEaten(i, eatens[i], wg)
	}
}

func (h Host) execApply(apply chan int, permits []chan bool, wg *sync.WaitGroup) {
	for {
		for id := range apply {
			//fmt.Printf("Apply from %v received\n", id)
			wg.Add(1)

			var MutexEatTime sync.Mutex
			MutexEatTime.Lock()

			permits[id] <- true

			MutexEatTime.Unlock()
			wg.Wait()
			//fmt.Printf("Wait done, last apply handle from %v\n", id)
		}
	}
}

func (h Host) execEaten(id int, eaten chan bool, wg *sync.WaitGroup) {
	for {
		<-eaten
		var MutexEatTime sync.Mutex
		MutexEatTime.Lock()
		h.eatTime[id]++
		MutexEatTime.Unlock()
		wg.Done()
	}
}

type Philo struct {
	id      int
	leftCS  *ChopS
	rightCS *ChopS
}

func (p Philo) start(permits []chan bool, eatens []chan bool, apply chan int, wgP *sync.WaitGroup) {
	go p.apply(apply, wgP)
	go p.eat(permits[p.id], eatens[p.id], wgP)
}

func (p Philo) apply(apply chan int, wgP *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		//fmt.Printf("Philo no. %v: Applying %vth time\n", p.id, i)
		wgP.Add(1)
		apply <- p.id
		//fmt.Printf("Philo no. %v: Success applied %vth time\n", p.id, i)
	}
}

func (p Philo) eat(permit chan bool, eaten chan bool, wgP *sync.WaitGroup) {
	var first *ChopS
	var second *ChopS
	if rand.Intn(2) == 0 {
		first = p.leftCS
		second = p.rightCS
	} else {
		first = p.rightCS
		second = p.leftCS
	}

	for ok := range permit {
		if ok {
			first.Lock()
			second.Lock()
			fmt.Printf("Philo no. %v: starting eating\n", p.id)
			first.Unlock()
			second.Unlock()
			fmt.Printf("Philo no. %v: finishing eating\n", p.id)
			eaten <- true
			wgP.Done()
			//fmt.Printf("Philo no. %v: true writed to eaten\n", p.id)
		}
	}
}

type ChopS struct {
	sync.Mutex
}
