// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Juju provider for CloudSigma

package xclarity

import (
	// "fmt"

	// "github.com/juju/errors"
	"github.com/juju/loggo"
	// "github.com/juju/utils"

	// "github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
)

var logger = loggo.GetLogger("juju.provider.xclarity")


type xclarityProvider struct {
	environCredentials
	environConfig
}

//********************************************
//
//	EnvironProvider interface
//  - config.Validator
//
//********************************************

func (xclarityProvider) Validate(cfg, oldCfg *config.Config) (*config.Config, error) {
	// Validate base configuration change before validating XClarity specifics.
	err := config.Validate(cfg, oldCfg)
	if err != nil {
		return nil, err
	}

	envCfg := &environConfig{
		Config: cfg,
		attrs: cfg.UnknownAttrs(),
	}
	return cfg.Apply(envCfg.attrs)
}

//********************************************
//
//	EnvironProvider interface
//  - Open
//
//  This interface is to initialize an Environ
//  which captures things needed to communicate
//  with a particular cloud, eg. ec2.
//********************************************
func (xclarityProvider) Open(params environs.OpenParams) (environs.Environ, error) {
	env := &xclarityEnviron{
		name: params.Config.Name(), 
		uuid: params.Config.UUID(), 
		config: *params.Config,
		cloudSpec: params.Cloud,
	}
	return *env, nil	
}