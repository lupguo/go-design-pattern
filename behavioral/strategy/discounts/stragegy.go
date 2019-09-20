/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package discounts

import "fmt"

type Strategy interface {
	Algorithm()
}

type StrategyCoupon struct {
	GoodsIds []int
	CouponId int
}

//Algorithm do coupon match algorithm
func (s *StrategyCoupon) Algorithm() {
	fmt.Println("Using coupon algorithm")
}

type StrategyFullMinus struct {
	GoodsIds []int
	total    float64
}

//Algorithm do full minus algorithm
func (s *StrategyFullMinus) Algorithm() {
	fmt.Println("Using Full minus algorithm")
}
