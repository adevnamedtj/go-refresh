package main

import "log"

func main() {

	seriesLength := 110
	fs := generate(seriesLength)
	log.Printf("fibonacci series for %v entries: %v, length: %v", seriesLength, fs, len(fs))
}

func generate(seriesLength int) (fs []int64) {
	previousVal := 0
	currentVal := 1

	fs = append(fs, int64(currentVal))
	for len(fs) < seriesLength {
		nextVal := previousVal + currentVal
		fs = append(fs, int64(nextVal))
		previousVal = currentVal
		currentVal = nextVal
	}

	return
}
