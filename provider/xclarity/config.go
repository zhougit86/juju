// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/schema"
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

// PrepareConfig is specified in the EnvironProvider interface.
// This is the first time during bootstrap that we will be touching configs. 
// Simply passed it on if cloud spec is valid. 
func (xclarityProvider) PrepareConfig(args environs.PrepareConfigParams) (*config.Config, error) {
	if err := validateCloudSpec(args.Cloud); err != nil {
		return nil, errors.Trace(err)
	}

	return args.Config, nil
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


var configFields = schema.Fields{}
var configDefaultFields = schema.Defaults{}
var configImmutableFields = []string{}

// This function will validate configuration values and populate defaults if necessary.
func validateConfig(cfg *config.Config, old *environConfig) (*environConfig, error) {
	// HOOK: where configurations are validated by underline cloud.

	// Check sanity of juju-level fields.
	var oldCfg *config.Config
	if old != nil {
		oldCfg = old.Config
	}
	if err := config.Validate(cfg, oldCfg); err != nil {
		return nil, errors.Trace(err)
	}

	// Extract validated provider-specific fields. All of configFields will be
	// present in validated, and defaults will be inserted if necessary. If the
	// schema you passed in doesn't quite express what you need, you can make
	// whatever checks you need here, before continuing.
	// In particular, if you want to extract (say) credentials from the user's
	// shell environment variables, you'll need to allow missing values to pass
	// through the schema by setting a value of schema.Omit in the configFields
	// map, and then to set and check them at this point. These values *must* be
	// stored in newAttrs: a Config will be generated on the user's machine only
	// to begin with, and will subsequently be used on a different machine that
	// will probably not have those variables set.
	newAttrs, err := cfg.ValidateUnknownAttrs(configFields, configDefaultFields)
	if err != nil {
		return nil, errors.Trace(err)
	}
	for field := range configFields {
		if newAttrs[field] == "" {
			return nil, errors.Errorf("%s: must not be empty", field)
		}
	}

	// If an old config was supplied, check any immutable fields have not changed.
	if old != nil {
		for _, field := range configImmutableFields {
			if old.attrs[field] != newAttrs[field] {
				return nil, errors.Errorf(
					"%s: cannot change from %v to %v",
					field, old.attrs[field], newAttrs[field],
				)
			}
		}
	}

	// Merge the validated provider-specific fields into the original config,
	// to ensure the object we return is internally consistent.
	newCfg, err := cfg.Apply(newAttrs)
	if err != nil {
		return nil, errors.Trace(err)
	}

	// Here we create the environConfig used throughout provider
	ecfg := &environConfig{
		Config: newCfg,
		attrs:  newAttrs,
	}

	return ecfg, nil
}