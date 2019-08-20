/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package socket

import (
	"errors"
	"fmt"
)

type Mode int

const (
	Ge2Cn Mode = iota
	Cn2Ge
)

type Socket interface {
	Insert()
	Provide()
}

//ge2cn the ability to switch between sockets(german standard => chinese standard)
type ge2cn struct {
}

func (a *ge2cn) Insert() {
	fmt.Println("Insert german standard socket")
}

func (a *ge2cn) Provide() {
	fmt.Println("Provide socket German => Chinese, now you can using [Chinese] electric products ")
}

//cn2ge the ability to switch between sockets(chinese standard => german standard)
type cn2ge struct {
}

func (a *cn2ge) Insert() {
	fmt.Println("Insert chinese standard socket")
}

func (a *cn2ge) Provide() {
	fmt.Println("Provide socket Chinese => German, now you can using [German] electric products")
}

type Adapter struct {
	ge2cn
}

//NewAdapter new an mode adapter
func NewAdapter(m Mode) (adapter Socket, err error) {
	switch m {
	case Cn2Ge:
		return new(cn2ge),nil
	case Ge2Cn:
		return new(ge2cn),nil
	}
	return nil, errors.New("have no such model adapter")
}
