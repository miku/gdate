package main

import (
	"time"

	"github.com/fatih/color"
)

func main() {
	t := time.Now().Format(time.RFC3339)
	color.Cyan(t)
}
