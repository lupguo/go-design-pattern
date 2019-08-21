# Object Pool Pattern

The object pool creational design pattern is used to prepare and keep multiple
instances according to the demand expectation.

> 简单的对象池（不包含外部资源，但只占用内存）可能效率不高，可能会降低性能. https://en.wikipedia.org/wiki/Object_pool_pattern

## Implementation

```go
package pool

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}
```

## Usage

Given below is a simple lifecycle example on an object pool.

```go
p := pool.New(2)

select {
case obj := <-p:
	obj.Do( /*...*/ )

	p <- obj
default:
	// No more objects left — retry later or fail
	return
}
```

## Rules of Thumb

- Object pool pattern is useful in cases where object initialization is more
  expensive than the object maintenance.
- If there are spikes in demand as opposed to a steady demand, the maintenance
  overhead might overweigh the benefits of an object pool.
- It has positive effects on performance due to objects being initialized beforehand.