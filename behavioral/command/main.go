/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import "github.com/tkstorm/go-design/behavioral/command/room"

func main() {
	// Receiver
	lamp := new(room.Light)

	// ConcreteCommand
	switchOn := &room.SwitchOnCommand{Light: lamp}
	switchOff := &room.SwitchOffCommand{Light: lamp}

	// Invoker
	mySwitch := new(room.Switch)
	mySwitch.StoreAndExecute(switchOn)
	mySwitch.StoreAndExecute(switchOff)
}
