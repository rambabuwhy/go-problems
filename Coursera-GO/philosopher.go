package main

import (
	"fmt"
	"sync"
	"time"
)

type chops struct {
	sync.Mutex
}

type philosopher struct {
	id    int
	left  *chops
	right *chops
}

var Wait sync.WaitGroup

func main() {

	count := 5

	sticks := make([]*chops, count)
	for i := 0; i < count; i++ {
		sticks[i] = new(chops)
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, left: sticks[i], right: sticks[(i+1)%count]}
		Wait.Add(1)
		go philosophers[i].eat()

	}
	Wait.Wait()
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.left.Lock()
		p.right.Lock()

		fmt.Printf("starting to eat: %d \n", p.id+1)
		time.Sleep(time.Second)

		p.right.Unlock()
		p.left.Unlock()

		fmt.Printf("finishing eating: %d \n", p.id+1)
		time.Sleep(time.Second)
	}
	Wait.Done()
}
