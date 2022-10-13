package main

// Restrict the number of executions of a critical tasks to a number

import (
	"fmt"
	"time"
)

var Token = struct{}{}

type Semaphore struct {
	tokens     chan struct{}
	total      int
	waitForAll chan struct{}
}

func NewSemaphore(capacity int, total int) *Semaphore {
	return &Semaphore{make(chan struct{}, capacity), total, make(chan struct{})}
}

func (semaphore *Semaphore) Acquire() {
	semaphore.tokens <- Token // try writing, is buffer is full, we'll block till somebody returns a token
}

func (semaphore *Semaphore) Release() {
	<-semaphore.tokens // drain a token, so somebody can put in a new token and start execution
	semaphore.total--
	if semaphore.total == 0 {
		close(semaphore.waitForAll)
	}
}

func (semaphore *Semaphore) WaitForAll() {
	<-semaphore.waitForAll
}

func main() {

	total := 25

	semaphore := NewSemaphore(4, total)

	for i := 0; i < total; i++ {
		go execute(semaphore, i) // create "total" number of goroutines, but only 4 would be active at any time
	}

	semaphore.WaitForAll()

}

func execute(semaphore *Semaphore, i int) {
	semaphore.Acquire()
	defer semaphore.Release()
	fmt.Println("Executing: ", i)
	time.Sleep(3 * time.Second)
}
