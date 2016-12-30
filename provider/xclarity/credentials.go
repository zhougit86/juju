// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
)

type environCredentials struct{}

const (
	credAttrUsername = "username"
	credAttrPassword = "password"
)

//********************************************
//
//	EnvironProvider interface
//  - ProviderCredentials
//
//********************************************

// CredentialSchemas is part of the environs.ProviderCredentials interface.
func (environCredentials) CredentialSchemas() map[cloud.AuthType]cloud.CredentialSchema {
	return map[cloud.AuthType]cloud.CredentialSchema{
		cloud.UserPassAuthType: {{
			"username", cloud.CredentialAttr{
				Description: "account username",
			},
		}, {
			"password", cloud.CredentialAttr{
				Description: "account password",
				Hidden:      true,
			},
		}},
	}
}

// DetectCredentials is part of the environs.ProviderCredentials interface.
func (environCredentials) DetectCredentials() (*cloud.CloudCredential, error) {
	return cloud.NewEmptyCloudCredential(), nil
}

// FinalizeCredential is part of the environs.ProviderCredentials interface.
func (environCredentials) FinalizeCredential(_ environs.FinalizeCredentialContext, args environs.FinalizeCredentialParams) (*cloud.Credential, error) {
	return &args.Credential, nil
}
