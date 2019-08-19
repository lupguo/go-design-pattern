package main

import "github.com/tkstorm/go-design/creational/factory/data"

func main() {
	// memory open
	memory, _ := data.NewStorage(data.MemoryStoreType)
	memory.Open("test.decolog")

	// disk open
	disk, _ := data.NewStorage(data.DiskStoreType)
	disk.Open("test.decolog")
}
