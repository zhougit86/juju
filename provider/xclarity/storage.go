// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/storage"
)

// StorageProviderTypes implements storage.ProviderRegistry.
func (xclarityEnviron) StorageProviderTypes() ([]storage.ProviderType, error) {
	return nil, nil
}

// StorageProvider implements storage.ProviderRegistry.
func (xclarityEnviron) StorageProvider(t storage.ProviderType) (storage.Provider, error) {
	return nil, errors.NotFoundf("storage provider %q", t)
}
