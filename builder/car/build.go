/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// car构建（颜色、轮子、速度）
package car

import "fmt"

type (
	Color string
	Wheels string
	Speed int
)

func (s Speed) String() string {
	return fmt.Sprintf("%d Km/s", s)
}

const (
	// 颜色
	ColorBlue  Color = "blue"
	ColorRed         = "red"
	ColorGreen       = "green"

	// 轮子
	WheelsFamily Wheels = "family wheels"
	WheelsSport         = "sport wheels"

	// 速度
	SpeedSportMax  Speed = 300
	SpeedFamilyMax       = 180
)

// 装配器
type Builder interface {
	// 喷漆
	paint(color Color)
	// 配轮胎
	assembleWheels(wheels Wheels)
	// 装引擎适配速度
	matchSpeed(speed Speed)
}

// 汽车
type Car struct {
	// car
	color    Color
	maxSpeed Speed
	wheels   Wheels

	curSpeed Speed
	working  bool
}

// build driver
func (c *Car) paint(color Color) {
	c.color = color
}

func (c *Car) assembleWheels(wheels Wheels) {
	c.wheels = wheels
}

func (c *Car) matchSpeed(speed Speed) {
	c.maxSpeed = speed
}

// 装配汽车
func Build(color Color, wheels Wheels, speed Speed) *Car {
	c := new(Car)
	c.paint(color)
	c.assembleWheels(wheels)
	c.matchSpeed(speed)
	return c
}
