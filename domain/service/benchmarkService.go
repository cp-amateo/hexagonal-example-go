package service

import (
	"math/rand"
	"time"
)

func simpleQuadraticAlgorithm(size int) int {
	rand.Seed(time.Now().Unix())
	array := rand.Perm(size)

	var result int

	for _, xValue := range array {
		for _, yValue := range array {
			xValue++
			yValue++
			result = xValue * yValue
			if result != 0 {
				result = (xValue - yValue) / result
			}
		}

	}

	return result

}
