package server

import (
	"log"
	"net"
	"sync"
)

// Job is a task which is submited by server to the worker pool
type Job struct {
	JobId int
	Conn  net.Conn
}

// WorkerPool is collection of thread responsible to do a certain task
type WorkerPool struct {
	MaxWorkers int             // MaxWorkers that worker pool will have
	QueueSize  int             // Number of task that will kept in queue if all workers are busy
	JobChan    chan Job        // Buffered channel to put in the job
	wg         *sync.WaitGroup // For uniformity across threads
}

// NewWorkerPool creates a pool of worker that procceses the request
func (w *WorkerPool) NewWorkerPool() {
	log.Println("Creating new worker pool")
	w.wg = new(sync.WaitGroup)
	w.JobChan = make(chan Job, w.MaxWorkers + w.QueueSize)
	for i:=0; i < w.MaxWorkers; i++ {
		w.wg.Add(1)
		log.Printf("Starting worker %d", i)
		go w.worker(i)
	}
}

// worker is thread which processes the request
func (w *WorkerPool) worker(workerId int) {
	processRequest := func (j Job) {
		request := make([]byte, 1024)
		j.Conn.Read(request)
		response := []byte("HTTP/1.1 200 OK\r\n\r\n Hello world ! \r\n")
		j.Conn.Write(response)
		j.Conn.Close()
	}
	for job := range w.JobChan {
		log.Printf("[worker %d] Processing request %d", workerId, job.JobId)
		processRequest(job)
	}
	w.wg.Done()
}

// SubmitJob puts the job into the channel and idle worker picks up
func (w* WorkerPool) SubmitJob(j Job) {
	w.JobChan <- j
}

// Close closes the channel and wait for all the workers to finish
func (w* WorkerPool) Close() {
	close(w.JobChan)
	w.wg.Wait()
}