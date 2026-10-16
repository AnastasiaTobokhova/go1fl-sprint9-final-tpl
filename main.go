package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000_000) + 1
	}
	return data
}

// maximum returns the maximum value, or zero for an empty slice.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks finds the maximum using CHUNKS goroutines.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxValues := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(index int, chunk []int) {
			defer wg.Done()
			if len(chunk) == 0 {
				maxValues[index] = data[0]
				return
			}
			maxValues[index] = maximum(chunk)
		}(i, data[start:end])
	}
	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
