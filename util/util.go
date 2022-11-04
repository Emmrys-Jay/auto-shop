package util

import (
	"fmt"
	"strings"
)

func Message(msg string) {
	fmt.Println()
	fmt.Print(strings.Repeat("~", 15), msg, strings.Repeat("~", 15), "\n")
	fmt.Println(strings.Repeat("~", 30+len(msg)))
}
