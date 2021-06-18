package gdate

import "time"

// Prefix adds a prefix.
func Prefix() string {
	return "Date: "
}

func Add(a, b int) int {
	return a + b
}

func AddSlow(a, b int) int {
	time.Sleep(1 * time.Millisecond)
	return a + b
}
