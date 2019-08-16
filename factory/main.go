package main

import "github.com/tkstorm/go-design/factory/data"

func main() {
	// memory open
	memory, _ := data.NewStorage(data.MemoryStoreType)
	memory.Open("test.log")

	// disk open
	disk, _ := data.NewStorage(data.DiskStoreType)
	disk.Open("test.log")
}
