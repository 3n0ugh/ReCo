package counters

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

func CountWord(fileName string) (int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, fmt.Errorf("file not found")
	}

	defer file.Close()

	// atomic counter variable
	var wordCount int64

	scanner := bufio.NewScanner(bufio.NewReader(file))
	// reading the file word by word
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		atomic.AddInt64(&wordCount, 1)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return wordCount, nil
}
