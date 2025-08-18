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
func levenshteinDistance(a, b string) int {
	la := len(a)
	lb := len(b)

	// Initialize matrix
	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}

	for i := 0; i <= la; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(
				dp[i-1][j]+1,      // deletion
				dp[i][j-1]+1,      // insertion
				dp[i-1][j-1]+cost, // substitution
			)
		}
	}

	return dp[la][lb]
}

// similarity returns a float64 between 0 and 1, where 1 means exact match
func similarity(a, b string) float64 {
	distance := levenshteinDistance(a, b)
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(distance)/float64(maxLen)
}

// IsSimilar returns true if input and target are similar above the given threshold
func IsSimilar(input, target string, threshold float64) bool {
	input = strings.ToLower(input)
	target = strings.ToLower(target)
	score := similarity(input, target)
	return score >= threshold
}
