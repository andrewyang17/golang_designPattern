package main

import "fmt"

// Least Frequently Used
type lfu struct {}

func (lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strategy")
}

