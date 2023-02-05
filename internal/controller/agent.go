package controller

import (
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/vo"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func ExecuteAgent(crd *crdtools.CRDConfig) error {
	crdContent, _ := crdtools.GetCRD(crd)
	dockerHost, _ := vo.TranslateToDockerHost(crdContent)
	fmt.Println(dockerHost)

	status := &docker.DockerHostStatus{
		Containers: []docker.DockerContainerSummary{},
		HostInfo: docker.DockerInfo{
			ID: "id",
		},
	}

	payload, err := vo.FromDockerHostStatus(status)
	if err != nil {
		fmt.Println("Error parsing to string")
		return err
	}
	crdtools.UpdateStatus(crd, payload)
	return nil
}
