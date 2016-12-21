// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/juju/environs"
)

const (
	providerType = "xclarity"
)

var providerInstance = xclarityProvider{}

func init() {
	environs.RegisterProvider(providerType, providerInstance)
}