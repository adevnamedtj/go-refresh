package main

import (
	"errors"
	"log"
	"math/rand"
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

	for i := 1; i < executions; i++ {
		inputArray, err := randomArray(10, i*10)

		if err != nil {
			log.Printf("Failed with %v", err)
			return
		}

		log.Printf("Input : %v", inputArray)

		outputArray := dedupInts(inputArray)
		log.Printf("Ouput : %v", outputArray)

	}

}

func dedupInts(values []int) []int {
	if len(values) == 0 {
		return values
	}

	uniqueValueKeyMap := make(map[int]int)
	duplicatedValueKeyMap := make(map[int][]int)
	var uniqueValues []int

	for i, v := range values {
		if _, ok := uniqueValueKeyMap[v]; !ok {
			uniqueValues = append(uniqueValues, v)
			uniqueValueKeyMap[v] = i
			continue
		}
		dupVal, okey := duplicatedValueKeyMap[v]

		if !okey {
			duplicatedValueKeyMap[v] = []int{i}
			continue
		}

		duplicatedValueKeyMap[v] = append(dupVal, i)

	}

	log.Printf("------------------------------------------------ Duplicates: %v", duplicatedValueKeyMap)

	return uniqueValues
}
