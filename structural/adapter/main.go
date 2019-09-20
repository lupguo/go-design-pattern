/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import (
	"github.com/tkstorm/go-design/structural/adapter/usb"
	"log"
)

type phone int

const (
	modeIphone phone = iota + 1
	modeAndroid
)

// 尝试为选择的手机型号，选择对应的USB数据线
// 	安卓和Iphone充电线都实现了Usb线的功能，即输入和输出（拥有适配功能）
func main() {
	var phoneMode phone = modeAndroid
	var line usb.Line
	switch phoneMode {
	case modeIphone:
		line = &usb.IphoneLine{}
	case modeAndroid:
		line = &usb.AndroidLine{}
	default:
		log.Fatal("not support phone mode")
	}

	line.Input()
	line.Output()

	//output:
	// common line to input
	// android line to output
}
