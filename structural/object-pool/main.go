package main

import (
	"context"
	"github.com/tkstorm/go-design/structural/object-pool/pool"
	"log"
	"time"
)

func main() {
	p := pool.New(5)
	for i := 0; i < 20; i++ {
		ctx, _ := context.WithTimeout(context.Background(), 200*time.Millisecond)
		select {
		case res := <-p:
			res.Do()
			// 资源归还
			p <- res
		case <-ctx.Done():
			log.Println("get resource timeout")
		}
	}
}
