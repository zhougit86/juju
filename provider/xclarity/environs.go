package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/constraints"
	"github.com/juju/juju/network"
	"github.com/juju/juju/storage"
)

// Environ is specific to each provider. 
// Here we define Environ for Lenovo XClarity.
type xclarityEnviron struct {
	name      string
	uuid      string
	config    config.Config
	cloudSpec environs.CloudSpec	
	host      string
}

//********************************************
//
//	Environ interface
//  - Provider
//
//********************************************
func (xclarityEnviron) Provider() environs.EnvironProvider {
	return providerInstance
}

//********************************************
//
//	InstanceBroker interfaces
//
//  This interface is part of the Environ
//  interface.  
//********************************************

func (e xclarityEnviron) AllInstances() ([]instance.Instance, error) {
	return e.Instances([]instance.Id{BootstrapInstanceId})
}

func (xclarityEnviron) StartInstance(args environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	return nil, errors.NotImplementedf("StartInstance")
}

func (xclarityEnviron) StopInstances(...instance.Id) error {
	return errors.NotImplementedf("StopInstance")
}

func (xclarityEnviron) MaintainInstance(args environs.StartInstanceParams) error {
	return errors.NotImplementedf("MaintainInstance")
}

//********************************************
//
//	Environ interface
//  - Instances
//
//  This interface returns a list of instances
//  based on instance ids.
//********************************************
func (e xclarityEnviron) Instances(ids []instance.Id) (instances []instance.Instance, err error) {
	instances = make([]instance.Instance, len(ids))
	var found bool
	for i, id := range ids {
		if id == BootstrapInstanceId {
			instances[i] = xclarityBootstrapInstance{e.host}
			found = true
		} else {
			err = environs.ErrPartialInstances
		}
	}
	if !found {
		err = environs.ErrNoInstances
	}
	return instances, err
}

//********************************************
//
//	Environ interface
//  - Bootstrap
//
//  Bootstrap creates a new environment, and an instance to host the
//  controller for that environment. The instnace will have have the
//  series and architecture of the Environ's choice, constrained to
//  those of the available tools. Bootstrap will return the instance's
//  architecture, series, and a function that must be called to finalize
//  the bootstrap process by transferring the tools and installing the
//  initial Juju controller.
//********************************************
func (xclarityEnviron) PrepareForBootstrap(ctx environs.BootstrapContext) error {
	return errors.NotImplementedf("PrepareForBootstrap")
}

func (xclarityEnviron) Bootstrap(
	ctx environs.BootstrapContext, 
	params environs.BootstrapParams,
) (*environs.BootstrapResult, error) {
	// Not implemented
	return nil, errors.NotImplementedf("Bootstrap")
}

func (xclarityEnviron) BootstrapMessage() string {
	return "hellow XClarity!"	
}

func (xclarityEnviron) Create(params environs.CreateParams) error {
	return errors.NotImplementedf("Create: "+params.ControllerUUID)
}

func (xclarityEnviron) ConstraintsValidator() (constraints.Validator, error) {
	return nil, errors.NotImplementedf("ConstraintsValidator")
}

func (xclarityEnviron) SetConfig(cfg *config.Config) error {
	return errors.NotImplementedf("SetConfig")
}

func (xclarityEnviron) ControllerInstances(controllerUUID string) ([]instance.Id, error) {
	return nil, errors.NotImplementedf("ControllerInstances")
}

func (xclarityEnviron) Destroy() error {
	return errors.NotImplementedf("Destroy")
}

func (xclarityEnviron) DestroyController(controllerUUID string) error {
	return errors.NotImplementedf("DestroyController")
}

func (xclarityEnviron) PrecheckInstance(series string, cons constraints.Value, placement string) error {
	return errors.NotImplementedf("PrecheckInstance")
}

//********************************************
//
//	Environ/Firewaller interface
//  - 
//********************************************

var errNoFwGlobal = errors.New("global firewall mode is not supported")

// OpenPorts is specified in the Environ interface. However, Azure does not
// support the global firewall mode.
func (xclarityEnviron) OpenPorts(ports []network.PortRange) error {
	return errNoFwGlobal
}

// ClosePorts is specified in the Environ interface. However, Azure does not
// support the global firewall mode.
func (xclarityEnviron) ClosePorts(ports []network.PortRange) error {
	return errNoFwGlobal
}

// Ports is specified in the Environ interface.
func (xclarityEnviron) Ports() ([]network.PortRange, error) {
	return nil, errNoFwGlobal
}

//********************************************
//
//	Environ/ConfigGetter interface
//  - 
//********************************************

func (e xclarityEnviron) Config() *config.Config {
	return &e.config
}

func (xclarityEnviron) StorageProviderTypes() ([]storage.ProviderType, error) {
	return nil, errors.NotImplementedf("StorageProviderTypes")
}

func (xclarityEnviron) StorageProvider(storage.ProviderType) (storage.Provider, error) {
	return nil, errors.NotImplementedf("StorageProvider")
}