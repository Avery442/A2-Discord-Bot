package src

import (
	"strings"
)

// min helper function
func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

// levenshteinDistance calculates the edit distance between two strings
// with memory efficiency (only two rows) and early exit when threshold is exceeded.
func levenshteinDistance(a, b string, maxDist int) int {
	la := len(a)
	lb := len(b)

	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	// Ensure a is the shorter string for memory efficiency
	if la > lb {
		a, b = b, a
		la, lb = lb, la
	}

	prevRow := make([]int, la+1)
	currRow := make([]int, la+1)

	for i := 0; i <= la; i++ {
		prevRow[i] = i
	}

	for j := 1; j <= lb; j++ {
		currRow[0] = j
		minInRow := currRow[0]

		for i := 1; i <= la; i++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			currRow[i] = min(
				prevRow[i]+1,      // deletion
				currRow[i-1]+1,    // insertion
				prevRow[i-1]+cost, // substitution
			)
			if currRow[i] < minInRow {
				minInRow = currRow[i]
			}
		}

		// Early exit: if minimum in row already exceeds maxDist
		if minInRow > maxDist {
			return maxDist + 1
		}

		// Swap rows
		prevRow, currRow = currRow, prevRow
	}

	return prevRow[la]
}

// similarity returns a float64 between 0 and 1, where 1 means exact match
func similarity(a, b string) float64 {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	if maxLen == 0 {
		return 1.0
	}

	// Set maxDist for early exit based on maxLen
	maxDist := maxLen
	distance := levenshteinDistance(a, b, maxDist)

	if distance > maxLen {
		return 0.0
	}

	return 1.0 - float64(distance)/float64(maxLen)
}

// IsSimilar returns true if input and target are similar above the given threshold
func IsSimilar(input, target string, threshold float64) bool {
	input = strings.ToLower(input)
	target = strings.ToLower(target)
	maxLen := len(input)
	if len(target) > maxLen {
		maxLen = len(target)
	}

	// Compute maximum allowed distance for early exit
	allowedDist := int(float64(maxLen) * (1.0 - threshold))
	distance := levenshteinDistance(input, target, allowedDist)

	if distance > allowedDist {
		return false
	}

	return true
}
