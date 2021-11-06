package pool

import (
	"errors"
	"log"
	"time"
)

// GoPool is to maintain a pool of Go routines to execute input tasks
type GoPool struct {
	// tasks are waiting to be executed.
	// when the channel is full, no more task is accepted.
	tasks chan func()
	// when a new value is pushed to chan, a new worker is spawned up.
	// when the channel is full, no more worker could be spawned.
	workers chan struct{}
}

func NewGoPool(taskQueueSize, maxWorkerQuantity, initWorkerQuantity int) *GoPool {
	pool := &GoPool{
		tasks:   make(chan func(), taskQueueSize),
		workers: make(chan struct{}, maxWorkerQuantity),
	}
	for i := 0; i < initWorkerQuantity; i++ {
		pool.workers <- struct{}{} // add a dummy data to channel to occupy the worker queue
		go pool.do(func() {})      // do some nonsen tasks here :D
	}
	return pool
}

// Queue puts the input task to pool to be executed.
func (g *GoPool) Queue(task func()) error {
	return g.queue(task, nil)
}

// QueueTimeout puts the input task to the pool,
// but the task only be executed in a specified duration.
func (g *GoPool) QueueTimeout(task func(), duration time.Duration) error {
	return g.queue(task, time.After(duration))
}

// queue puts the input task to the pool.
func (g *GoPool) queue(task func(), timeout <-chan time.Time) error {
	select {
	case <-timeout:
		return errors.New("Timeout to do the task")
	case g.tasks <- task: // task queue is not full, so the task will be done by some running go routine
		log.Println("add to task queue")
		return nil
	case g.workers <- struct{}{}: // task queue is full, spawns a new go routine to do the task
		log.Println("spawns new routine")
		go g.do(task)
		return nil
	}
}

// do gonna executes the input task,
// if the task queue is not empty, it also executes those tasks
func (g *GoPool) do(task func()) {
	defer func() { <-g.workers }()
	task() // do assigned task
	for task := range g.tasks {
		task() // do the remaining tasks in the queue
	}
}
