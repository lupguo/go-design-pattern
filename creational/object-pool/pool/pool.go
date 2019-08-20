/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// 资源池模型，是一种资源预支，资源复用，适合稳定业务的处理模式：
//	- 对象资源复用：对象池模式在对象初始化很重的情况下非常有用。
//	- 适合稳定业务：如果需求出现峰值而不是稳定的需求，则开销可能会超过对象池的好处。
//	- 资源预支：由于对象是预先初始化的，它对性能有积极的影响。
package pool

import (
	"log"
	"math/rand"
	"time"
)

// 资源
type Resource struct {
	No int
}

// 资源干事(随机做0~400ms)
func (r *Resource) Do() {
	time.Sleep(time.Duration(rand.Intn(5)) * 100 * time.Millisecond)
	log.Printf("resource #%d work finish\n", r.No)
}

// 资源池(基于通道)
type Pool chan *Resource

// New 创建指定大小的资源池（对一些初始化很重的资源很有用，比如TCP连接、Redis、Mysql认证等）
func New(size int) Pool {
	p := make(Pool, size)
	for i := 0; i < size; i++ {
		// slow create
		time.Sleep(500*time.Millisecond)
		p <- &Resource{No: i}
	}
	return p
}

