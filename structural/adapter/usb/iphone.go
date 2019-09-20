/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package usb

import "fmt"

type IphoneLine struct {
	length int
}

func (l *IphoneLine) Input() {
	fmt.Println("common line to input")
}

func (l *IphoneLine) Output() {
	fmt.Println("iphone line to output")
}



