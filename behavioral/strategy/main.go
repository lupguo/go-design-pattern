/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import (
	"github.com/tkstorm/go-design/behavioral/strategy/discounts"
	"math/rand"
)

// 策略模式，将算法封装在单独的对象中。类将算法委托给对象，而不是直接实现算法
func main() {
	// ... 分析订单商品清单，检测满足优惠策略情况（此处将具体的优惠券、满减算法被屏蔽）
	var strategy discounts.Strategy
	couponId := rand.Intn(1)
	goodsIds := []int{80, 81, 82}
	switch {
	case couponId > 5:
		strategy = &discounts.StrategyCoupon{
			GoodsIds: goodsIds,
			CouponId: couponId,
		}
	default:
		strategy = &discounts.StrategyFullMinus{
			GoodsIds: goodsIds,
		}
	}
	strategy.Algorithm()
}
