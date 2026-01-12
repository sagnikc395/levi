package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
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

func main() {
	// // Define flags
	// srcPtr := flag.String("src", "", "The source string.")
	// destPtr := flag.String("dest", "", "The destination string.")

	// flag.Parse()

	// // Print ASCII Art
	// fmt.Printf("\033[32m%s\033[0m\n", ASCII)

	// // Show help if both options are empty (matching the Python logic)
	// if *srcPtr == "" && *destPtr == "" {
	// 	flag.Usage()
	// 	os.Exit(0)
	// }

	// // Convert strings to runes to handle Unicode correctly and slice efficiently
	// s1 := []rune(*srcPtr)
	// s2 := []rune(*destPtr)

	// distance := lev(s1, s2)

	// fmt.Printf("\033[32m>>> Levenshtein distance bw '%s' and '%s' is %d\033[0m\n", *srcPtr, *destPtr, distance)

	cmd := &cli.Command{
		Name:  "levi",
		Usage: "compare levenshtein distance between strings or files",
		Action: func(context.Context, *cli.Command) error {
			fmt.Println("levi is here")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
