package pool

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

type TaskType int

const (
	ReadFromSocket TaskType = iota
	SomethingElse
)

type TaskFD struct {
	task func()
	fd   int
	t    TaskType
}

// GoPool is to maintain a pool of Go routines to execute input tasks
type GoPool struct {
	// tasks are waiting to be executed.
	// when the channel is full, no more task is accepted.
	tasks chan TaskFD
	// when a new value is pushed to chan, a new worker is spawned up.
	// when the channel is full, no more worker could be spawned.
	workers chan struct{}
	// which file descriptors are under processing
	fd map[int]map[TaskType]bool

	lock *sync.Mutex
}

func (g *GoPool) Status(trim bool) string {
	if trim {
		if len(g.tasks) > 0 || len(g.workers) > 1 {
			return fmt.Sprintf("tasks: %v, workers: %v", len(g.tasks), len(g.workers))
		}
	}
	return fmt.Sprintf("tasks: %v, workers: %v", len(g.tasks), len(g.workers))
}

func NewGoPool(taskQueueSize, maxWorkerQuantity, initWorkerQuantity int) *GoPool {
	log.Printf(
		"create pool: queue: %v, max worker: %v, init worker: %v",
		taskQueueSize, maxWorkerQuantity, initWorkerQuantity,
	)
	pool := &GoPool{
		tasks:   make(chan TaskFD, taskQueueSize),
		workers: make(chan struct{}, maxWorkerQuantity),
		fd:      make(map[int]map[TaskType]bool),
		lock:    &sync.Mutex{},
	}
	for i := 0; i < initWorkerQuantity; i++ {
		pool.workers <- struct{}{} // add a dummy data to channel to occupy the worker queue
		go pool.do(TaskFD{
			task: func() {
				log.Println("initial worker :D")
			},
			fd: -1,
			t:  SomethingElse,
		}) // do some nonsen tasks here :D
	}
	return pool
}

// Queue puts the input task to pool to be executed.
func (g *GoPool) Queue(task func(), fd int, t TaskType) error {
	log.Println("4. ", g.Status(false))
	g.lock.Lock()
	defer g.lock.Unlock()
	if g.fd[fd] == nil {
		g.fd[fd] = map[TaskType]bool{}
	}
	if g.fd[fd][t] == false {
		g.fd[fd][t] = true
		return g.queue(TaskFD{task: task, fd: fd, t: t}, nil)
	} else {
		log.Println("duplicate task for fd: ", fd, t)
		return nil
	}
}

// QueueTimeout puts the input task to the pool,
// but the task only be executed in a specified duration.
func (g *GoPool) QueueTimeout(task func(), fd int, duration time.Duration) error {
	return g.queue(TaskFD{task: task, fd: fd}, time.After(duration))
}

// queue puts the input task to the pool.
func (g *GoPool) queue(task TaskFD, timeout <-chan time.Time) error {
	for {
		select {
		case <-timeout:
			log.Printf("5. worker status: task: %v, worker: %v", len(g.tasks), len(g.workers))
			return errors.New("Timeout to do the task")
		case g.tasks <- task: // task queue is not full, so the task will be done by some running go routine
			log.Printf("5. worker status: task: %v, worker: %v", len(g.tasks), len(g.workers))
			log.Println("6. add to task queue")
			return nil
		// case g.workers <- struct{}{}: // task queue is full, spawns a new go routine to do the task
		// log.Printf("pool status: task: %v, worker: %v", len(g.tasks), len(g.workers))
		// 	log.Println("spawns new routine")
		// 	go g.do(task)
		// 	return nil
		default:
			log.Println("why it go here?")
			return nil
		}
	}
}

// do gonna executes the input task,
// if the task queue is not empty, it also executes those tasks
func (g *GoPool) do(task TaskFD) {
	defer func() {
		v, ok := <-g.workers
		log.Println("stop worker: ", v, ok)
		log.Printf("pool status: task: %v, worker: %v", len(g.tasks), len(g.workers))
	}()
	task.task() // do assigned task
	g.lock.Lock()
	if g.fd[task.fd] == nil {
		g.fd[task.fd] = make(map[TaskType]bool)
	}
	g.fd[task.fd][task.t] = false
	g.lock.Unlock()
	for task := range g.tasks {
		log.Println("fd to: ", g.fd)
		task.task() // do the remaining tasks in the queue
		g.lock.Lock()
		g.fd[task.fd][task.t] = false
		g.lock.Unlock()
	}
}
