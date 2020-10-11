package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	workers := 3
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var jobResults []int

	inputChan := make(chan int, len(jobs))
	outputChan := make(chan int, len(jobs))

	// initiating woker pool
	initaiatePool(workers, inputChan, outputChan)

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

func initaiatePool(workers int, requestChannel chan int, reponseChannel chan int) {
	for i := 1; i <= workers; i++ {
		go worker(i, requestChannel, reponseChannel)
	}
}

func worker(workerID int, requestChannel chan int, reponseChannel chan int) {
	for job := range requestChannel {
		log.Printf("worker %v processing job %v", workerID, job)
		reponseChannel <- run(job)
	}
}

func run(job int) int {
	// job impl
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return job
}
