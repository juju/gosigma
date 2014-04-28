// Copyright 2014 ALTOROS
// Licensed under the AGPLv3, see LICENSE file for details.

package gosigma

import (
	"fmt"

	"github.com/Altoros/gosigma/data"
)

// A RuntimeNIC represents runtime information for network interface card
type RuntimeNIC struct {
	obj *data.RuntimeNetwork
}

// Type of network interface card (public, private, etc)
func (r RuntimeNIC) Type() string {
	if r.obj != nil {
		return r.obj.InterfaceType
	} else {
		return ""
	}
}

// AddressIPv4 returns runtime IPv4 address (if any)
func (r RuntimeNIC) AddressIPv4() string {
	if r.obj != nil && r.obj.IPv4 != nil {
		return r.obj.IPv4.UUID
	} else {
		return ""
	}
}

// String method is used to print values passed as an operand to any format that
// accepts a string or to an unformatted printer such as Print.
func (r RuntimeNIC) String() string {
	return fmt.Sprintf(`{Type: %q, Address: %q}`, r.Type(), r.AddressIPv4())
}
