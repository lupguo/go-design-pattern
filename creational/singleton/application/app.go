/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// 单例创建过程中，需要考虑资源竟态问题
package application

import (
	"fmt"
	"sync"
)

// App单例
type App map[string]interface{}

var (
	once sync.Once
	app  App
)

// 创建App单例，需要考虑竟态问题
func New() App {
	once.Do(func() {
		app = make(App)
	})
	fmt.Printf("pointer = %#p\n", app)
	return app
}

// Config App配置
func (app App) Config(key string, val interface{}) {
	app[key] = val
}

// GetConfig 获取配置
func (app App) GetConfig(key string) interface{} {
	if v, ok := app[key]; ok {
		return v
	}
	return nil
}
