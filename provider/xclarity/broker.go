package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
	"github.com/juju/juju/storage"
)

//********************************************
//
//	InstanceBroker interfaces
//
//  This interface is part of the Environ
//  interface.  
//********************************************

func (env xclarityEnviron) AllInstances() ([]instance.Instance, error) {
	return env.Instances([]instance.Id{BootstrapInstanceId})
}


// This is where we build an instance with user input parameters and constraints.
// It is the integration point where juju meets underline cloud provider.
// For XClarity, I'm assuming that a REST post with parameters, constraints will be sent
// to XClarity, and response with information that to populate StartInstanceResult.
func (env xclarityEnviron) StartInstance(args environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	hardware := instance.HardwareCharacteristics{}
    volumes := make([]storage.Volume, 0)
	networkInfo := make([]network.InterfaceInfo, 0)
	volumeAttachments := make([]storage.VolumeAttachment, 0)

	return &environs.StartInstanceResult{
		Instance:          xclarityBootstrapInstance{},
		Config:			   &env.config,
		Hardware:          &hardware, // type instance.HardwareCharacteristics struct
		NetworkInfo:       networkInfo, // type network.InterfaceInfo struct
		Volumes:           volumes, // type storage.Volume struct
		VolumeAttachments: volumeAttachments, // type storageVolumeAttachment struct
	}, nil
}

func (xclarityEnviron) StopInstances(...instance.Id) error {
	return errors.NotImplementedf("StopInstance")
}

func (xclarityEnviron) MaintainInstance(args environs.StartInstanceParams) error {
	return errors.NotImplementedf("MaintainInstance")
}