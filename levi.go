package main

// lev calculates the Levenshtein distance between two rune slices recursively.
func levi(s1, s2 []rune) int {
	// Case 1: s1 is empty
	if len(s1) == 0 {
		return len(s2)
	}
	// Case 2: s2 is empty
	if len(s2) == 0 {
		return len(s1)
	}

	// Case if the last characters match
	if s1[len(s1)-1] == s2[len(s2)-1] {
		return levi(s1[:len(s1)-1], s2[:len(s2)-1])
	}

	// Recursive calls for Insert, Remove, and Replace
	return 1 + min(
		levi(s1[:len(s1)-1], s2),             // Deletion
		levi(s1, s2[:len(s2)-1]),             // Insertion
		levi(s1[:len(s1)-1], s2[:len(s2)-1]), // Substitution
	)
}
