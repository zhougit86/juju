package xclarity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/juju/errors"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/instance"
	"github.com/juju/juju/network"
	"github.com/juju/juju/storage"
	"github.com/juju/utils/arch"
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

type DummyREST struct {
	Birthday string `json:"birthday"`
	Name     string `json:"eng_name"`
}

func TryRESTTest(id string) (*DummyREST, error) {

	logger.Debugf("feng TryRESTTest")

	url := fmt.Sprintf("http://192.168.5.11:8000/api/v1/person/%s/?format=json&charset=utf8", id)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Errorf("NewRequest: ", err)
		return nil, err
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Send REST query: ", err)
		return nil, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record DummyREST

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		logger.Errorf("JSON decoder error:", err)
		return nil, err
	}

	logger.Infof("Birthday = ", record.Birthday)
	logger.Infof("Name   = ", record.Name)
	return &record, nil
}

// This is where we build an instance with user input parameters and constraints.
// It is the integration point where juju meets underline cloud provider.
// For XClarity, I'm assuming that a REST post with parameters, constraints will be sent
// to XClarity, and response with information that to populate StartInstanceResult.
func (env xclarityEnviron) StartInstance(args environs.StartInstanceParams) (*environs.StartInstanceResult, error) {
	// HOOK: start an instance through XClarity, and acquire responses that can populate the following struct
	// for juju to register and manage this intance.
	//
	// For now, we are imitating a successful instance by directly returning a result struct
	var tmpArch string = arch.AMD64
	var tmpMem uint64 = 2000000
	var tmpCpuCore uint64 = 1
	var tmpCpuPower uint64 = 100

	logger.Debugf("feng Deploy 4 devx", args)

	hardware := instance.HardwareCharacteristics{
		Arch:     &tmpArch,
		Mem:      &tmpMem,
		CpuCores: &tmpCpuCore,
		CpuPower: &tmpCpuPower,
	}
	volumes := make([]storage.Volume, 0)
	networkInfo := make([]network.InterfaceInfo, 0)
	volumeAttachments := make([]storage.VolumeAttachment, 0)

	TryRESTTest("1")

	return &environs.StartInstanceResult{
		Instance:          xclarityBootstrapInstance{},
		Config:            env.ecfg.Config,
		Hardware:          &hardware,         // type instance.HardwareCharacteristics struct
		NetworkInfo:       networkInfo,       // type network.InterfaceInfo struct
		Volumes:           volumes,           // type storage.Volume struct
		VolumeAttachments: volumeAttachments, // type storageVolumeAttachment struct
	}, nil
}

func (xclarityEnviron) StopInstances(...instance.Id) error {
	// HOOK: Ask XClarity to stop instance.
	return nil
}

func (xclarityEnviron) MaintainInstance(args environs.StartInstanceParams) error {
	return errors.NotImplementedf("MaintainInstance")
}
