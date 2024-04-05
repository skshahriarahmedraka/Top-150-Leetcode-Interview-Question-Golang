package main


// func minLevelsToGainMorePoints(possible []int) int {
func minimumLevels(possible []int) int {
   n := len(possible)
	// Calculate total points for Danielchandg and Bob if they play optimally
	totalPoints := 0
	for _, val := range possible {
		totalPoints += val
	}
	if totalPoints <= 0 {
		return -1 // If total points are non-positive, Danielchandg can't gain more points
	}

	// Calculate prefix sums
	prefixSums := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSums[i+1] = prefixSums[i] + possible[i]
	}

	// Calculate the minimum number of levels Danielchandg should play to gain more points
	minLevels := n + 1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			danielPoints := prefixSums[j] - prefixSums[i]
			bobPoints := totalPoints - danielPoints
			if danielPoints > bobPoints && j-i < minLevels {
				minLevels = j - i
			}
		}
	}

	if minLevels == n+1 {
		return -1
	}
	return minLevels
}