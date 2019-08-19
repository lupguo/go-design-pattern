/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// output:
// 2019/08/19 23:43:03 Welcome to seven day shop!
// 2019/08/19 23:43:03 Thank you for your purchase!
// 2019/08/19 23:43:03 Goodbye and welcome to come again!
package main

import "github.com/tkstorm/go-design/behavioral/observer/shop"

func main() {
	// create an seven day shop
	sevenDay := shop.NewSevenDay()

	// register guard as observer
	guard := new(shop.Guard)
	sevenDay.Register(guard)

	// event notifier
	sevenDay.Notifier(shop.EventCome)
	sevenDay.Notifier(shop.EventBuy)
	sevenDay.Notifier(shop.EventTalk) // talk event not care
	sevenDay.Notifier(shop.EventLeave)

	// remove guard observer
	sevenDay.Deregister(guard)

	// now event no report
	sevenDay.Notifier(shop.EventCome)
	sevenDay.Notifier(shop.EventBuy)
	sevenDay.Notifier(shop.EventTalk)
	sevenDay.Notifier(shop.EventLeave)
}
