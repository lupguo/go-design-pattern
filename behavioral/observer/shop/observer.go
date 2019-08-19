/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package shop

//EventType 事件类型
type EventType int

//Event 监听事件
type Event struct {
	eventId EventType
	desc    string
}

//NewEvent 创建事件类型
func NewEvent(eventId EventType, desc string) *Event {
	return &Event{
		eventId: eventId,
		desc:    desc,
	}
}

//Observer 事件观察者
type Observer interface {
	// 观察者在关注事件
	OnNotify(e *Event)
}

//Notifier 被观察者观察内容
type Notifier interface {
	Register(o Observer)
	Deregister(o Observer)
	Notifier(e *Event)
}
