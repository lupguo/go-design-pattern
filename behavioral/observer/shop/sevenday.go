/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package shop

import (
	"log"
)

// 7天店，事件相关类型
const (
	Access EventType = 10 * (1 + iota)
	Buy
	Leave
	Talk
)

var (
	EventCome  = NewEvent(Access, "customer coming")
	EventBuy   = NewEvent(Buy, "customer buying")
	EventLeave = NewEvent(Leave, "customer leaving")
	EventTalk  = NewEvent(Talk, "talk with sales")
)

//SevenDay 被观察内容
type SevenDay struct {
	observerList []Observer
}

//NewSevenDay 创建7天门店
func NewSevenDay() *SevenDay {
	return &SevenDay{
	}
}

//Register 注册门店观察者
func (sd *SevenDay) Register(o Observer) {
	sd.observerList = append(sd.observerList, o)
}

//Deregister 移除门店观察者
func (sd *SevenDay) Deregister(o Observer) {
	for k, v := range sd.observerList {
		if v == o {
			sd.observerList[k] = nil
		}
	}
}

//Notifier 绑定门店监控指定事件
func (sd *SevenDay) Notifier(e *Event) {
	for _, o := range sd.observerList {
		if o == nil {
			continue
		}
		o.OnNotify(e)
	}
}

//Guard 7天店，门禁部分
type Guard struct {
}

//OnNotify 观察者通知事件/发布事件
func (b *Guard) OnNotify(e *Event) {
	switch e.eventId {
	case Access:
		log.Println("Welcome to seven day shop!")
	case Buy:
		log.Println("Thank you for your purchase!")
	case Leave:
		log.Println("Goodbye and welcome to come again!")
	}
}
