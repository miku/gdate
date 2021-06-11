package main

import (
	"time"

	"github.com/fatih/color"
	"github.com/miku/gdate"
)

func main() {
	t := time.Now().Format(time.RFC3339)
	color.Green(gdate.Preamble() + t)
}
