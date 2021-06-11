package main

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/miku/gdate"
	"go4.org/strutil"
)

func main() {
	t := time.Now().Format(time.RFC3339)
	if strutil.IsPlausibleJSON(t) {
		log.Println("that's weird")
	}
	color.Green(gdate.Preamble() + t)
}
