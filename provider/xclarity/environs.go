package xclarity

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/juju/errors"
	"github.com/juju/juju/constraints"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
	"github.com/juju/juju/provider/common"
	"github.com/juju/utils/arch"
)

// Environ is specific to each provider.
// Here we define Environ for Lenovo XClarity.
type xclarityEnviron struct {
	mu        sync.Mutex
	name      string
	uuid      string
	ecfg      *environConfig
	cloudSpec environs.CloudSpec
	host      string
}

var _ environs.Environ = (*xclarityEnviron)(nil)

//********************************************
//
//	Environ interface
//  - Provider
//
//********************************************
func (*xclarityEnviron) Provider() environs.EnvironProvider {
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
func (e *xclarityEnviron) Instances(ids []instance.Id) (instances []instance.Instance, err error) {
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
func (*xclarityEnviron) PrepareForBootstrap(ctx environs.BootstrapContext) error {
	// If nothing, return nil
	return nil
}

func (env *xclarityEnviron) Bootstrap(
	ctx environs.BootstrapContext,
	params environs.BootstrapParams,
) (*environs.BootstrapResult, error) {

	// result := environs.BootstrapResult{
	// 	Arch: arch.AMD64,
	// 	Series: "boostraped series",
	// 	Finalize: nil,
	// }
	return common.Bootstrap(ctx, env, params)
	// return &result, nil
}

func (*xclarityEnviron) BootstrapMessage() string {
	return "xClarity bootstraped! hello world."
}

func (*xclarityEnviron) Create(params environs.CreateParams) error {
	// lxd, cloudsigma provider
	return nil
}

// Interface function where we define what type of vocabulary of constraints
// that can be taken by XClarity cloud. For example,
func (env *xclarityEnviron) ConstraintsValidator() (constraints.Validator, error) {
	validator := constraints.NewValidator()

	// Register unsupported constraints
	validator.RegisterUnsupported([]string{
		constraints.Container,    // do not support container, yet
		constraints.InstanceType, // do not support instance type
		constraints.Tags,         // do not support tagging
		constraints.VirtType,     // do not support multi-hypervisor
	})

	// Register constraints that XClarity cloud can support
	validator.RegisterVocabulary(constraints.Arch, []string{arch.AMD64})

	return validator, nil
}

// Interface function used to initialize/update environConfig.
func (env *xclarityEnviron) SetConfig(cfg *config.Config) error {
	env.mu.Lock()
	defer env.mu.Unlock()

	// Save old if any
	var old *environConfig
	if env.ecfg != nil {
		old = env.ecfg
	}

	// Validate and populate with defaults
	// so the returned configs are valid for consumption
	environConfig, err := validateConfig(cfg, old)
	if err != nil {
		return errors.Trace(err)
	}

	// Set the environment value
	env.ecfg = environConfig

	// Done
	return nil
}

func (*xclarityEnviron) ControllerInstances(controllerUUID string) ([]instance.Id, error) {
	return nil, errors.NotImplementedf("ControllerInstances")
}

func (*xclarityEnviron) Destroy() error {
	return errors.NotImplementedf("Destroy")
}

func (env *xclarityEnviron) DestroyController(controllerUUID string) error {
	return common.Destroy(env)
}

func (*xclarityEnviron) PrecheckInstance(series string, cons constraints.Value, placement string) error {
	// HOOK: This is called in "juju deploy".
	// Can ask XClarity for verifications.
	return nil
}

//********************************************
//
//	Environ/Firewaller interface
//  -
//********************************************

var errNoFwGlobal = errors.New("global firewall mode is not supported")

// OpenPorts is specified in the Environ interface. However, Azure does not
// support the global firewall mode.
func (*xclarityEnviron) OpenPorts(ports []network.PortRange) error {
	return errNoFwGlobal
}

// ClosePorts is specified in the Environ interface. However, Azure does not
// support the global firewall mode.
func (*xclarityEnviron) ClosePorts(ports []network.PortRange) error {
	return errNoFwGlobal
}

// Ports is specified in the Environ interface.
func (*xclarityEnviron) Ports() ([]network.PortRange, error) {
	return nil, errNoFwGlobal
}

//********************************************
//
//	Environ/ConfigGetter interface
//  -
//********************************************
func (env *xclarityEnviron) Config() *config.Config {
	// For read, do not need lock protection.
	return env.ecfg.Config
}

func mytrace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}

func mycaller() {
	_, file, no, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("called from %s#%d\n", file, no)
	}
}
