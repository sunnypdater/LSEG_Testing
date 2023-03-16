package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define a Task struct to hold task and callback
type Task struct {
    task     int
    callback func(error, interface{})
}

// Define a Queue struct
type Queue struct {
    workerFunc  func(int) (interface{}, error)
    taskChannel chan Task
    paused      bool
    mutex       sync.Mutex
}

// Worker function to execute tasks
func (q *Queue) worker() {
    for {
        // Wait for task in taskChannel
        task := <-q.taskChannel

        // Execute task asynchronously
        res, err := q.workerFunc(task.task)

        // Execute the callback function
        task.callback(err, res)
    }
}

// Push a task to the queue
func (q *Queue) Push(task int, callback func(error, interface{})) {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    // Create a Task object with task and callback
    t := Task{
        task:     task,
        callback: callback,
    }

    // Push the task to the taskChannel
    q.taskChannel <- t
}

// Pause the execution of tasks
func (q *Queue) Pause() {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    q.paused = true
}

// Resume the execution of tasks
func (q *Queue) Resume() {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    q.paused = false
}

// Return the current execution state
func (q *Queue) Paused() bool {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    return q.paused
}

// Return the number of tasks in the queue
func (q *Queue) Length() int {
    q.mutex.Lock()
    defer q.mutex.Unlock()

    return len(q.taskChannel)
}

// Initialize a new Queue object
func NewQueue(workerFunc func(int) (interface{}, error)) *Queue {
    q := &Queue{
        workerFunc:  workerFunc,
        taskChannel: make(chan Task),
        paused:      false,
    }

    // Start the worker function in a goroutine
    go q.worker()

    return q
}

// Worker function that waits a random amount of time then returns the task x2
func Worker(task int) (interface{}, error) {
    randomDelay := time.Duration(time.Duration(1000) * time.Millisecond * time.Duration(time.Duration(rand.Intn(1000))))
    time.Sleep(randomDelay)

    return task * 2, nil
}

// Callback function to handle task results or errors
func Callback(err error, res interface{}) {
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(res)
    }
}

func main() {
    q := NewQueue(Worker)

    q.Push(1, Callback)
    q.Push(2, Callback)
    q.Push(3, Callback)

    // Responses:
    // 2
    // 4
    // 6

    q.Pause()
    fmt.Println(q.Paused()) // true

    q.Push(4, Callback)
    q.Push(5, Callback)

    fmt.Println(q.Length()) // 2

    q.Resume()

    // Responses:
    // 8
    // 10
}