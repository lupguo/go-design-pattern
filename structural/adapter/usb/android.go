/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package usb

import "fmt"

type Line interface {
	Input()
	Output()
}

type AndroidLine struct {
	length int
}

func (l *AndroidLine) Input() {
	fmt.Println("common line to input")
}

func (l *AndroidLine) Output() {
	fmt.Println("android line to output")
}

//Input 适配Usb输入
func Input(l Line)  {
	l.Input()
}

//Output 适配Usb输出
func Output(l Line)  {
	l.Output()
}