// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

// Juju provider for CloudSigma

package xclarity

import (
	// "fmt"

	// "github.com/juju/errors"
	// "github.com/juju/loggo"
	// "github.com/juju/utils"

	// "github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
)

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

func (p xclarityProvider) Validate(cfg, oldCfg *config.Config) (*config.Config, error) {
	logger.Infof("+++++++++++++++ 1 +++++++++++++++++++++")

	// Validate base configuration change before validating XClarity specifics.
	err := config.Validate(cfg, oldCfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
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
	env := xclarityEnviron{
		name: params.Config.Name(), 
		uuid: params.Config.UUID(), 
		cloudSpec: params.Cloud,
	}

	// Set environConfig
	environConfig, err := validateConfig(params.Config, nil)
	if err != nil {
		logger.Errorf("xclarity.provider.Open", err)
		return nil, err
	}

	// Set the environment value
	env.ecfg = environConfig

	// TODO: why cannt I use env.SetConfig!? Don't understand.
	// if err := env.SetConfig(params.Config); err != nil {
	// 	return nil, err
	// }

	// Environment is complete
	return env, nil	
}