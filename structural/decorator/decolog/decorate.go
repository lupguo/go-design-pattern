/*
 * Copyright (c) 2019 TerryRod <tkstorm1988@gmail.com>
 *
 * Except as contained in this notice, the name(s) of the above copyright holders
 * shall not be used in advertising or otherwise to promote the sale, use or other
 * dealings in this Software without prior written authorization.
 */

package decolog

import (
	"log"
	"time"
)

//OperateFn represents operations that require decoration
type OperateFn func()

//Decorate the operation
func Decorate(opFn OperateFn)  {
	defer func(s time.Time) {
		log.Printf("elpased time %0.2d ms", time.Since(s).Nanoseconds()/(1<<20))
	}(time.Now())

	// real operation function
	opFn()
}