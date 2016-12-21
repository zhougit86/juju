package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/status"
	"github.com/juju/juju/network"
)

const (
	// BootstrapInstanceId is the instance ID used
	// for the manual provider's bootstrap instance.
	BootstrapInstanceId instance.Id = "xclarity:"
)

type xclarityBootstrapInstance struct {
	host string
}

func (xclarityBootstrapInstance) Id() instance.Id {
	return BootstrapInstanceId
}

func (xclarityBootstrapInstance) Status() instance.InstanceStatus {
	// We asume that if we are deploying in manual provider the
	// underlying machine is clearly running.
	return instance.InstanceStatus{
		Status: status.Running,
	}
}

func (xclarityBootstrapInstance) Refresh() error {
	return errors.NotImplementedf("Refresh")
}

func (inst xclarityBootstrapInstance) Addresses() (addresses []network.Address, err error) {
	// Not implemented
	return nil, errors.NotImplementedf("Addresses")
}

func (xclarityBootstrapInstance) OpenPorts(machineId string, ports []network.PortRange) error {
	return errors.NotImplementedf("OpenPorts")
}

func (xclarityBootstrapInstance) ClosePorts(machineId string, ports []network.PortRange) error {
	return errors.NotImplementedf("ClosePorts")
}

func (xclarityBootstrapInstance) Ports(machineId string) ([]network.PortRange, error) {
	return nil, errors.NotImplementedf("Ports")
}