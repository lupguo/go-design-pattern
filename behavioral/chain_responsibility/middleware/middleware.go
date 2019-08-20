/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

// 为解除请求的发送者和接收者之间耦合，而使多个对象都有机会处理这个请求。
// 将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它。
package middleware

import (
	"fmt"
	"net/http"
)

type Handler interface {
	// Next handler
	Next(handler Handler) Handler

	// Handle http request
	Handle(req *http.Request)
}

//NextHandler transmit set handler to next handler
type NextHandler struct {
	next Handler
}

func (h *NextHandler) Next(handler Handler) Handler {
	h.next = handler
	return h
}

func (h *NextHandler) Handle(req *http.Request) {
	fmt.Println("handle begin")

	// next handle
	if h.next != nil {
		h.next.Handle(req)
	}
}

//BlackUaHandler block black ua
type BlackUaHandler struct {
	NextHandler
}

func (h *BlackUaHandler) Handle(req *http.Request) {
	fmt.Println("analyze useragent...")

	// next handle
	if h.next != nil {
		h.next.Handle(req)
	}
}

//BlackIpHandler
type BlackIpHandler struct {
	NextHandler
}

func (h *BlackIpHandler) Handle(req *http.Request) {
	fmt.Println("analyze request ip...")

	// next handle
	if h.next != nil {
		h.next.Handle(req)
	}
}

type CommonHttpHandler struct {
	NextHandler
}

func (h CommonHttpHandler) Handle(req *http.Request) {
	fmt.Println("deal with http request and return with response!")
}
