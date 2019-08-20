/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package main

import "github.com/tkstorm/go-design/structural/adapter/socket"

func main()  {
	// german socket to chinese socket
	g2c, _ := socket.NewAdapter(socket.Ge2Cn)
	g2c.Insert()
	g2c.Provide()

	// chinese socket to german socket
	c2g, _ := socket.NewAdapter(socket.Cn2Ge)
	c2g.Insert()
	c2g.Provide()
}
