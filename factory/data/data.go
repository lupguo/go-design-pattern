// 工厂模式
package data

import (
	"errors"
	"fmt"
)

type Store interface {
	Open(file string)
}

type StoreType int

const (
	MemoryStoreType = iota + 1
	DiskStoreType
	TempStoreType
)

// 内存存储
type memoryStore struct {
	Name string
	Size int
}

func (m *memoryStore) Open(file string) {
	fmt.Println("# memory storage open")
}

// 磁盘存储
type diskStore struct {
	Name string
	Size int
}

func (f *diskStore) Open(file string) {
	fmt.Println("# disk storage open")
}

// 临时文件存储
type tempStore struct {
	Name string
	size int
}

func (t *tempStore) Open(file string) {
	fmt.Println("# template storage open")
}

// 工厂方式创建存储实例
func NewStorage(t StoreType) (s Store, err error) {
	switch t {
	case MemoryStoreType:
		return &memoryStore{"mem storage", 1 << 6}, nil
	case DiskStoreType:
		return &diskStore{"disk storage", 1 << 9}, nil
	case TempStoreType:
		return &tempStore{"temp storage", 1 << 6}, nil
	default:
		return nil, errors.New("storage types not yet supported")
	}
}
