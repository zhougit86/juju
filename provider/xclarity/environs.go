package xclarity

import (
	"sync"
	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/constraints"
	"github.com/juju/juju/network"
	"github.com/juju/juju/provider/common"
	"github.com/juju/juju/instance"
)

// Environ is specific to each provider. 
// Here we define Environ for Lenovo XClarity.
type xclarityEnviron struct {
	mu        sync.Mutex
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
	logger.Debugf("xclarity PrepareForBootstrap")

	// If nothing, return nil
	return nil
}

func (env xclarityEnviron) Bootstrap(
	ctx environs.BootstrapContext, 
	params environs.BootstrapParams,
) (*environs.BootstrapResult, error) {
	return common.Bootstrap(ctx, env, params)
}

func (xclarityEnviron) BootstrapMessage() string {
	return "hellow XClarity!"	
}

func (xclarityEnviron) Create(params environs.CreateParams) error {
	return errors.NotImplementedf("Create: "+params.ControllerUUID)
}

// Borrowed from cloudsigma/environcaps.go
var unsupportedConstraints = []string{
	constraints.Container,
	constraints.InstanceType,
	constraints.Tags,
	constraints.VirtType,
}

func (xclarityEnviron) ConstraintsValidator() (constraints.Validator, error) {
	validator := constraints.NewValidator()
	validator.RegisterUnsupported(unsupportedConstraints)
	return validator, nil
}

func (env xclarityEnviron) SetConfig(cfg *config.Config) error {
	env.mu.Lock()
	defer env.mu.Unlock()

	var old config.Config
	if &env.config != nil {
		old = env.config
	}
	ecfg, err := providerInstance.Validate(cfg, &old)
	if err != nil {
		return err
	}
	env.config = *ecfg

	return nil
}

func (xclarityEnviron) ControllerInstances(controllerUUID string) ([]instance.Id, error) {
	return nil, errors.NotImplementedf("ControllerInstances")
}

func (xclarityEnviron) Destroy() error {
	return errors.NotImplementedf("Destroy")
}

func (env xclarityEnviron) DestroyController(controllerUUID string) error {
	return common.Destroy(env)
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