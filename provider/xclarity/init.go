// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/loggo"
	"github.com/juju/juju/environs"
)

var logger = loggo.GetLogger("juju.provider.xclarity")

const (
	providerType = "xclarity"
)

var providerInstance = xclarityProvider{}

func init() {
	environs.RegisterProvider(providerType, providerInstance)
}