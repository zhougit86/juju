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

// This is a top-level entry point used by juju's CLI
var providerInstance = xclarityProvider{}

func init() {
	environs.RegisterProvider(providerType, providerInstance)
}