package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func maxMTCharacters(n, k int, s string) int {
	// Count the frequency of each character
	frequency := make(map[rune]int)
	for _, char := range s {
		frequency[char]++
	}

	// Create a slice of frequencies
	var freqSlice []int
	for _, freq := range frequency {
		freqSlice = append(freqSlice, freq)
	}

	// Sort the slice in decreasing order
	sort.Sort(sort.Reverse(sort.IntSlice(freqSlice)))

	// We want to maximize the number of 'M' and 'T', so we will add operations to the max of M or T
	// If we have more 'M' or the same amount of 'M' and 'T', we focus on 'M', else 'T'
	mCount := frequency['M']
	tCount := frequency['T']

	// Assign k operations to the maximum of 'M' or 'T'
	if mCount >= tCount {
		mCount += k
	} else {
		tCount += k
	}

	// The maximum number is the sum of 'M' and 'T' after operations
	return mCount + tCount
}

func main() {
	// Reading input
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)
	s, _ := reader.ReadString('\n')

	// Calculate and print the result
	result := maxMTCharacters(n, k, s)
	fmt.Println(result)
}
