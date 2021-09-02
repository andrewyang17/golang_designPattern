package main

import "fmt"

// First In First Out
type fifo struct {}

func (fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strategy")
}
