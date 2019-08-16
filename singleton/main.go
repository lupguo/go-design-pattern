package main

import (
	"fmt"
	"github.com/tkstorm/go-design/singleton/application"
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
	fmt.Println("app1 host is", app.GetConfig("host"))
	fmt.Println("app2 host is", app.GetConfig("host"))
}
