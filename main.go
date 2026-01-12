package main

import (
	"flag"
	"fmt"
	"os"
)

const ASCII string = `
                          _____                                     
 _____               _____\    \ _______    ______   ____________   
|\    \             /    / |    |\      |  |      | /            \  
 \\    \           /    /  /___/| |     /  /     /||\___/\  \\___/| 
  \\    \         |    |__ |___|/ |\    \  \    |/  \|____\  \___|/ 
   \|    | ______ |       \       \ \    \ |    |         |  |      
    |    |/      \|     __/ __     \|     \|    |    __  /   / __   
    /            ||\    \  /  \     |\         /|   /  \/   /_/  |  
   /_____/\_____/|| \____\/    |    | \_______/ |  |____________/|  
  |      | |    ||| |    |____/|     \ |     | /   |           | /  
  |______|/|____|/ \|____|   | |      \|_____|/    |___________|/   
                         |___|/                                    

a levenshtein distance comparision utility tool
`

// lev calculates the Levenshtein distance between two rune slices recursively.
func lev(s1, s2 []rune) int {
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
		return lev(s1[:len(s1)-1], s2[:len(s2)-1])
	}

	// Recursive calls for Insert, Remove, and Replace
	return 1 + min(
		lev(s1[:len(s1)-1], s2),             // Deletion
		lev(s1, s2[:len(s2)-1]),             // Insertion
		lev(s1[:len(s1)-1], s2[:len(s2)-1]), // Substitution
	)
}

// min helper function for integers
func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func main() {
	// Define flags
	srcPtr := flag.String("src", "", "The source string.")
	destPtr := flag.String("dest", "", "The destination string.")

	flag.Parse()

	// Print ASCII Art
	fmt.Printf("\033[32m%s\033[0m\n", ASCII)

	// Show help if both options are empty (matching the Python logic)
	if *srcPtr == "" && *destPtr == "" {
		flag.Usage()
		os.Exit(0)
	}

	// Convert strings to runes to handle Unicode correctly and slice efficiently
	s1 := []rune(*srcPtr)
	s2 := []rune(*destPtr)

	distance := lev(s1, s2)

	fmt.Printf("\033[32m>>> Levenshtein distance bw '%s' and '%s' is %d\033[0m\n", *srcPtr, *destPtr, distance)
}
