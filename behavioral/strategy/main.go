/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import "github.com/tkstorm/go-design/behavioral/strategy/brake"

// 策略行为设计模式允许在运行时选择算法的行为。
// The policy behavior design pattern allows the behavior of the algorithm to be selected at run time.
func main() {

	type Car struct {
		brakeSystem brake.System
	}

	suvCar := &Car{brake.NewAbsBrake("ABS BRAKE 2.0")}
	truckCar := &Car{brake.NewAbsBrake("NORMAL BRAKE 1.8")}

	suvCar.brakeSystem.Brake()
	truckCar.brakeSystem.Brake()
}
