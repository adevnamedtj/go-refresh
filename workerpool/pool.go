package main

import (
	"log"
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
		go func(workerID int) {
			for job := range requestChannel {
				log.Printf("worker %v processing job %v", workerID, job)
				reponseChannel <- runJob(job)
			}
		}(i)
	}
}

func runJob(j int) int {
	time.Sleep(1 * time.Second)
	return j
}
