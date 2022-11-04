package util

import (
	"fmt"
	"io"
	"strings"
)

func Message(w io.Writer, msg string) {
	fmt.Fprintln(w)
	fmt.Fprint(w, strings.Repeat("~", 15), msg, strings.Repeat("~", 15), "\n")
	fmt.Fprintln(w, strings.Repeat("~", 30+len(msg)))
}
