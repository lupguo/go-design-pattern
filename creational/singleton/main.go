/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// 确保一个类只有一个实例，并提供对该实例的全局访问。
package main

import (
	"fmt"
	"github.com/tkstorm/go-design/creational/singleton/application"
)

func main() {
	// App为map，复制基于指针Copy
	app := application.New()
	app1 := application.New()
	app2 := application.New()
	fmt.Printf("%#p, %#p, %#p, %#p\n", application.New(), &app, &app1, &app2)

	// set config
	app.Config("host", "https://tkstorm.com")
	fmt.Println("app host is", app.GetConfig("host"))
	fmt.Println("app1 host is", app1.GetConfig("host"))
	fmt.Println("app2 host is", app2.GetConfig("host"))
}
