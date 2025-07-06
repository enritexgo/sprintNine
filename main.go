package main

import (
	"fmt"
	"math/rand"
	"slices"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	for i := range result {
		result[i] = rand.Int()
	}

	return result
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	return slices.Max(data)

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	var wg sync.WaitGroup
	size := len(data) / CHUNKS
	maxCounts := make([]int, CHUNKS)

	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		sliceFirst := i * size
		sliceEnd := sliceFirst + size

		go func() {
			defer wg.Done()
			maxCounts[i] = slices.Max(data[sliceFirst:sliceEnd])
		}()
	}
	wg.Wait()
	return slices.Max(maxCounts)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	var (
		start   time.Time
		max     int
		elapsed time.Duration
	)

	start = time.Now()
	max = maximum(data)
	elapsed = time.Duration(time.Since(start).Microseconds())

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Duration(time.Since(start).Microseconds())
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
