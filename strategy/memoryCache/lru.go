package main

import "fmt"

// Least Recently Used
type lru struct {}

func (lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

