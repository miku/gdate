package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/miku/gdate"
)

func main() {
	msg := fmt.Sprintf("%s %s\n", gdate.Prefix(), time.Now().Format(time.RFC3339))
	color.Cyan(msg)
}
