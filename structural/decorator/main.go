/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// 修饰模式是类继承的另外一种选择，可以通过装饰模式来做到，在运行时给类增加行为。
//
// 类继承在编译时候增加行为，而装饰模式是在运行时增加行为。(go中没有类的概念，但可以通过接口实现做到行为增加）
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
	// decorate log a
	decolog.Decorate(decolog.ActionFn(DoActionA))

	// decorate log b
	decolog.Decorate(decolog.ActionFn(DoActionB))
}

func DoActionA() {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	log.Println("finish do action a")
}

func DoActionB() {
	time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	log.Println("finish do action b")
}
