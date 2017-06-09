/*-
 * Copyright © 2017, Jörg Pernfuß <code.jpe@gmail.com>
 * All rights reserved.
 *
 * Use of this source code is governed by a 2-clause BSD license
 * that can be found in the LICENSE file.
 */

package twister // import "github.com/mjolnir42/twister/lib/twister"

import (
	"runtime"

	"github.com/mjolnir42/legacy"
)

// Dispatch implements erebos.Dispatcher
func Dispatch(msg []byte) error {
	// send all messages from the same host to the same
	// handler to keep the ordering intact
	hostID, err := legacy.PeekHostID(msg)
	if err != nil {
		return err
	}

	Handlers[hostID%runtime.NumCPU()].InputChannel() <- msg
	return nil
}

// vim: ts=4 sw=4 sts=4 noet fenc=utf-8 ffs=unix
