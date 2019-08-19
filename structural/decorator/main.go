/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// run result
// 2019/08/19 19:05:24 finish action a
//2019/08/19 19:05:24 elpased time 77 ms
//2019/08/19 19:05:24 finish action b
//2019/08/19 19:05:24 elpased time 88 ms

package main

import (
	"github.com/tkstorm/go-design/structural/decorator/decolog"
	"log"
	"math/rand"
	"time"
)

func main() {
	// convert to decolog.ActionFn
	fna := decolog.ActionFn(DoActionA)
	fnb := decolog.ActionFn(DoActionB)

	// decorate log a
	decolog.Decorate(fna)

	// decorate log b
	decolog.Decorate(fnb)
}

func DoActionA()  {
	time.Sleep(time.Duration(rand.Intn(200))*time.Millisecond)
	log.Println("finish action a")
}

func DoActionB()  {
	time.Sleep(time.Duration(rand.Intn(200))*time.Millisecond)
	log.Println("finish action b")
}


