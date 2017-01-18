package xclarity

import (
	"github.com/juju/errors"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
	"github.com/juju/juju/status"
)

const (
	// BootstrapInstanceId is the instance ID used
	// for the manual provider's bootstrap instance.
	BootstrapInstanceId instance.Id = "xclarity: "
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
	// HOOK: this is where xclarity needs to tell me the address of bootstraped machine 0

	// For now, we are giving it a known IP that can function as machine 0.
	newAddress := network.NewAddress("192.168.8.234")
	aaah := []network.Address{newAddress}
	return aaah, nil
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
