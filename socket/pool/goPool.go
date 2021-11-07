package pool

import (
	"errors"
	"fmt"
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

func (g *GoPool) Status() string {
	return fmt.Sprintf("tasks: %v, workers: %v", len(g.tasks), len(g.workers))
}

func NewGoPool(taskQueueSize, maxWorkerQuantity, initWorkerQuantity int) *GoPool {
	log.Printf(
		"create pool: queue: %v, max worker: %v, init worker: %v",
		taskQueueSize, maxWorkerQuantity, initWorkerQuantity,
	)
	pool := &GoPool{
		tasks:   make(chan func(), taskQueueSize),
		workers: make(chan struct{}, maxWorkerQuantity),
	}
	for i := 0; i < initWorkerQuantity; i++ {
		pool.workers <- struct{}{} // add a dummy data to channel to occupy the worker queue
		go pool.do(func() {
			log.Println("initial worker :D")
		}) // do some nonsen tasks here :D
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
	log.Printf("pool status: task: %v, worker: %v", len(g.tasks), len(g.workers))
	select {
	case <-timeout:
		return errors.New("Timeout to do the task")
	// case g.tasks <- task: // task queue is not full, so the task will be done by some running go routine
	// 	log.Println("add to task queue")
	// 	return nil
	case g.workers <- struct{}{}: // task queue is full, spawns a new go routine to do the task
		log.Println("spawns new routine")
		go g.do(task)
		return nil
	default:
		log.Println("why it go here?")
		return nil
	}
}

// do gonna executes the input task,
// if the task queue is not empty, it also executes those tasks
func (g *GoPool) do(task func()) {
	defer func() {
		v, ok := <-g.workers
		log.Println("stop worker: ", v, ok)
		log.Printf("pool status: task: %v, worker: %v", len(g.tasks), len(g.workers))
	}()
	task() // do assigned task
	// for task := range g.tasks {
	// 	task() // do the remaining tasks in the queue
	// }
}
