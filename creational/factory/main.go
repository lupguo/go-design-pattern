/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

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
