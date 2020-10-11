package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	jobTimeOut := 2 * time.Second
	workers := 3
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var jobResults []int

	inputChan := make(chan int, len(jobs))
	outputChan := make(chan int, len(jobs))

	// initiating woker pool
	initaiatePool(workers, jobTimeOut, inputChan, outputChan)

	// feeding the pool
	for _, job := range jobs {
		inputChan <- job
	}

	// getting results
	for i := 1; i <= len(jobs); i++ {
		result := <-outputChan
		jobResults = append(jobResults, result)
	}

	close(inputChan)
	close(outputChan)

	log.Printf("Input: %v", jobs)
	log.Printf("Ouput: %v", jobResults)

}

func initaiatePool(workers int, jobTimeOut time.Duration, requestChannel, reponseChannel chan int) {
	for i := 1; i <= workers; i++ {
		go worker(i, jobTimeOut, requestChannel, reponseChannel)
	}
}

func worker(workerID int, jobTimeOut time.Duration, requestChannel, reponseChannel chan int) {
	for job := range requestChannel {
		log.Printf("worker %v processing job %v", workerID, job)
		reponseChannel <- run(job, jobTimeOut)
	}
}

func run(job int, jobTimeOut time.Duration) int {
	rChan := make(chan int, 1)
	go callDownstream(job, rChan)
	select {
	case r := <-rChan:
		return r
	case <-time.After(jobTimeOut):
		close(rChan)
		log.Printf("Downstream timeout for job: %v", job)
		return 0
	}
}

func callDownstream(job int, rChan chan int) {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	rChan <- job
}
