/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package car

import (
	"fmt"
	"strings"
)

const FireInitSpeed Speed = 20

// Work
type Work interface {
	// 点火
	Fire()
	// 停止
	Stop()
	// 加速
	SpeedUp(s Speed)
	// 减速
	SpeedDown(s Speed)
}

// driver attribute show
func (c *Car) attrShow() string {
	return fmt.Sprintf("Car Information: Color=>%s, MaxSpeed=>%d, Wheels:%s, CurrentSpeed:%d",
		c.color, c.maxSpeed, c.wheels, c.curSpeed)
}

// driver message show
func (c *Car) message(info ...interface{}) {
	var sct []string
	for _, v := range info {
		switch vv := v.(type) {
		case string:
			sct = append(sct, vv)
		default:
			sct = append(sct, fmt.Sprint(v))
		}
	}
	if sct != nil {
		fmt.Printf("%[2]s\n\t %[1]s\n", c.attrShow(), strings.Join(sct, " "))
	} else {
		fmt.Println(c.attrShow())
	}
}

// drive driver
func (c *Car) Fire() {
	if !c.working {
		c.working = true
		c.curSpeed = FireInitSpeed
		c.message("fire")
	} else {
		c.message("already fired")
	}
}

func (c *Car) Stop() {
	c.working = false
	c.curSpeed = 0
	c.message("stop the driver")
}

func (c *Car) SpeedUp(s Speed) {
	if c.working {
		// speed up
		switch {
		case s > c.maxSpeed:
			c.curSpeed = c.maxSpeed
			c.message("speed to max speed", c.curSpeed)
		case s > FireInitSpeed && s <= c.maxSpeed:
			c.curSpeed = s
			c.message("speed to", c.curSpeed)
		default:
			c.message("pls step on the gas")
		}
	}
}

func (c *Car) SpeedDown(s Speed) {
	if c.working {
		// speed down
		switch {
		case s <= 0:
			c.Stop()
		case s < c.curSpeed:
			c.curSpeed = s
			c.message("slow speed to", c.curSpeed)
		}
	}
}
