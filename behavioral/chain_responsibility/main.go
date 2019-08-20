/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import "github.com/tkstorm/go-design/behavioral/chain_responsibility/middleware"

// output:
// handle begin
// analyze useragent...
// analyze request ip...
// deal with http request and return with response!
func main() {
	// handler
	handler := new(middleware.NextHandler)
	uaHandler := new(middleware.BlackUaHandler)
	ipHandler := new(middleware.BlackIpHandler)
	cmHandler := new(middleware.CommonHttpHandler)

	// chain of responsibility
	handler.Next(uaHandler)
	uaHandler.Next(ipHandler)
	ipHandler.Next(cmHandler)

	// handler handle
	handler.Handle(nil)
}
