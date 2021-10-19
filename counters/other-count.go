package counters

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

func CountOthers(fileName string) ([]int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("file not found")
	}

	defer file.Close()

	// atomic counter variables
	var vowelCount int64
	var consonantCount int64
	var pMarkCount int64
	var digitCount int64
	var spaceCount int64

	jobs := make(chan string, 20)
	done := make(chan bool)

	go func() {
		// reading the file line by line
		scanner := bufio.NewScanner(bufio.NewReader(file))
		for scanner.Scan() {
			jobs <- scanner.Text() // sending lines on jobs channel
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
		close(jobs)
	}()
	// starting goroutine
	go countQuery(jobs, done, &vowelCount, &pMarkCount, &digitCount, &consonantCount, &spaceCount)

	// waiting to end of countQuery
	<-done
	close(done)

	// adding results to result slice
	var result []int64
	result = append(result, vowelCount, consonantCount, spaceCount, digitCount, pMarkCount, vowelCount+consonantCount)
	return result, nil
}

func countQuery(jobs <-chan string, done chan<- bool, vowelCount, pMarkCount, digitCount, consonantCount, spaceCount *int64) {

	for w := range jobs { // receiving words from jobs channel
		for _, s := range w { // making conditions on words
			switch s {
			// vowels ascii values
			case 97, 101, 105, 111, 117, 65, 69, 73, 79, 85:
				atomic.AddInt64(vowelCount, 1)
			// punctation marks ascii values
			case 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
				46, 47, 58, 59, 60, 61, 62, 63, 64, 92, 93, 95, 96,
				123, 124, 125, 126, 8217:
				atomic.AddInt64(pMarkCount, 1)
			// digits ascii values
			case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
				atomic.AddInt64(digitCount, 1)
			case 32, 10:
				atomic.AddInt64(spaceCount, 1)
			default:
				atomic.AddInt64(consonantCount, 1)
			}
		}
	}

	done <- true
}
