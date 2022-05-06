package utils

import (
	"fmt"
	"strings"
	"time"
)

func TypeText(text string, interval time.Duration) {
	fmt.Println("\x1b[2J\x1b[H") // Clear screen
	lines := strings.Split(text, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		fmt.Print("  ")
		for j := 0; j < len(line); j++ {
			fmt.Print(string(line[j]))
			time.Sleep(interval)
		}
		fmt.Println()
	}
	fmt.Println()
	time.Sleep(2 * time.Second)
}
