/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package room

import "fmt"

//light - Receiver
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("The light is on")
}

func (l *Light) TurnOff() {
	fmt.Println("The light is off")
}

//Command - Command Interface
type Command interface {
	execute()
}

//SwitchOnCommand - ConcreteCommand(The Command for turning on the room)
type SwitchOnCommand struct {
	Light *Light
}

func (cmd *SwitchOnCommand) execute() {
	cmd.Light.TurnOn()
}

//SwitchOffCommand - ConcreteCommand(The Command for turning off the room)
type SwitchOffCommand struct {
	Light *Light
}

func (cmd *SwitchOffCommand) execute() {
	cmd.Light.TurnOff()
}

//Switch - The invoker struct
type Switch struct {
	commands []Command
}

//StoreAndExecute
func (sw *Switch) StoreAndExecute(cmd Command)  {
	sw.commands = append(sw.commands, cmd)
	cmd.execute()
}
