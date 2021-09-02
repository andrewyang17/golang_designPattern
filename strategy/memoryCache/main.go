package main

func main() {
	lfu := &lfu{}
	lru := &lru{}
	fifo := &fifo{}

	cache := newCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	cache.setEvictionAlgo(lru)
	cache.add("d", "3")

	cache.setEvictionAlgo(fifo)
	cache.add("e", "3")
}