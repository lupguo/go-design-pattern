/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package brake

import "fmt"

const (
	AbsBrakeType = 1 << iota
	NorBrakeType
)

//System brake system
type System interface {
	Brake()
}

//AbsBrake abs brake system
type AbsBrake struct {
	mode string
}

func NewAbsBrake(mode string) *AbsBrake {
	return &AbsBrake{mode: mode}
}

func (b *AbsBrake) Brake() {
	fmt.Println(b.mode, "using abs brake.")
}

//NormalBrake
type NormalBrake struct {
	mode string
}

func NewNormalBrake(mode string) *NormalBrake {
	return &NormalBrake{mode: mode}
}

func (b *NormalBrake) Brake() {
	fmt.Println(b.mode, "using normal brake.")
}
