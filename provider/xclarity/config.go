// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
)

type environConfig struct {
	*config.Config
	attrs map[string]interface{}
}

//********************************************
//
//	EnvironProvider interface
//  - PrepareConfig
//
//  This interface captures actions to manage
//  config values, things like validation,
//  CRUD, etc.
//********************************************

// Borrowed from provider/maas.go
// PrepareConfig is specified in the EnvironProvider interface.
func (p xclarityProvider) PrepareConfig(args environs.PrepareConfigParams) (*config.Config, error) {
	if err := validateCloudSpec(args.Cloud); err != nil {
		return nil, errors.Trace(err)
	}
	envConfig, err := p.Validate(args.Config, nil)
	if err != nil {
		return nil, err
	}
	//return args.Config.Apply(environConfig.attrs)
	return envConfig, nil
}

// Borrowed from maas/provider.go
func validateCloudSpec(spec environs.CloudSpec) error {
	if spec.Endpoint == "" {
		return errors.Errorf(
			"missing address of host to bootstrap: " +
				`please specify "juju bootstrap xclarity/[user@]<host>"`,
		)
	}
	return nil
}