/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// Builder通过分离构造和表示来构造复杂对象
// 这里通过car.build(喷漆、轮子、引擎)来构建一台赛车或者私家车
package main

import (
	"fmt"
	"github.com/tkstorm/go-design/creational/builder/car"
)

func main() {
	fmt.Println("Build race car...")

	// in car design race car
	raceCar := car.Build(car.ColorBlue, car.WheelsSport, car.SpeedSportMax)
	raceCar.Fire()
	// speed up
	raceCar.SpeedUp(100)
	raceCar.SpeedUp(400)
	// speed down
	raceCar.SpeedDown(300)
	raceCar.SpeedDown(200)
	raceCar.SpeedDown(5)
	raceCar.SpeedDown(0)

	fmt.Println("Build family car...")
	familyCar := car.Build(car.ColorGreen, car.WheelsFamily, car.SpeedFamilyMax)
	familyCar.Fire()
	familyCar.SpeedUp(300)
	familyCar.Stop()
}