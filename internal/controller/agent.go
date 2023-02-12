package controller

import (
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/dockerutils"
	"github.com/6zacode-toolbox/docker-agent/internal/vo"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func ExecuteAgent(crd *crdtools.CRDConfig) error {
	crdContent, _ := crdtools.GetCRD(crd)
	dockerHost, _ := vo.TranslateToDockerHost(crdContent)
	fmt.Println(dockerHost)

	dockerInfo, _ := dockerutils.ExecuteDockerInfo()
	dockerPS, _ := dockerutils.ExecuteDockerPS()

	status := &docker.DockerHostStatus{
		Containers: dockerPS,
		DockerHost: dockerInfo,
	}

	payload, err := vo.FromDockerHostStatus(status)
	if err != nil {
		fmt.Println("Error parsing to string")
		return err
	}
	crdtools.UpdateStatus(crd, payload)
	return nil
}

func ExecuteDockerComposeRunner(crd *crdtools.CRDConfig) error {
	crdContent, _ := crdtools.GetCRD(crd)
	composeRunner, _ := vo.TranslateToDockerComposeRunner(crdContent)
	fmt.Println(composeRunner)

	// Do something
	// Execute Logic
	dockerRunner, err := dockerutils.ExecuteCompose(&composeRunner)
	if err != nil {
		fmt.Println("Error executing command")
		return err
	}
	fmt.Println(dockerRunner)
	status := &docker.DockerComposeRunnerStatus{}

	payload, err := vo.FromDockerComposeRunnerStatus(status)
	if err != nil {
		fmt.Println("Error parsing to string")
		return err
	}
	crdtools.UpdateStatus(crd, payload)
	return nil
}
