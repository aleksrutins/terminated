package utils

import (
	"fmt"
	"time"
)

func TypeText(text string, interval time.Duration) {
	fmt.Print("\x1b[2J\x1b[H") // Clear screen
	for i := 0; i < len(text); i++ {
		fmt.Print(string(text[i]))
		time.Sleep(interval)
	}
	fmt.Println()
	time.Sleep(2 * time.Second)
}
