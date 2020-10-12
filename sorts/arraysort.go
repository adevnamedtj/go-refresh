package main

import (
	"errors"
	"log"
	"math/rand"
	"sort"
	"time"
)

func randomArray(length int, maxInteger int) (randAry []int, err error) {
	if length == 0 || maxInteger == 0 {
		return randAry, errors.New("length or maxIntegr can not be zero")
	}
	rand.Seed(time.Now().UnixNano())

	for i := 1; i < length; i++ {
		randAry = append(randAry, rand.Intn(maxInteger))
	}

	return
}

func main() {

	executions := 10
	assendingOrder := false

	for i := 1; i < executions; i++ {
		inputArray, err := randomArray(10, i*10)

		if err != nil {
			log.Printf("Failed with %v", err)
			return
		}

		log.Printf("Input : %v", inputArray)

		outputArray := sortByOrder(assendingOrder, inputArray)
		log.Printf("Ouput : %v", outputArray)

	}

}

func sortByOrder(assendingOrder bool, ints []int) []int {

	if assendingOrder {

		sort.Slice(ints, func(i, j int) bool {
			return ints[i] < ints[j]
		})

		return ints
	}

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] > ints[j]
	})

	return ints

}
