package vo

import (
	"encoding/json"

	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

type WithStatus struct {
	Status docker.DockerHostStatus `json:"status,omitempty"`
}

func TranslateToDockerComposeRunner(payload []byte) (docker.DockerComposeRunner, error) {
	blankObject := docker.DockerComposeRunner{}
	resultObject := docker.DockerComposeRunner{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		//client.Logger.Info(fmt.Sprintf("Error loading CRD   #%v ", err))
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}

func FromDockerComposeRunnerStatus(object *docker.DockerComposeRunnerStatus) (string, error) {
	type WithStatus struct {
		Status docker.DockerComposeRunnerStatus `json:"status,omitempty"`
	}
	withStatus := WithStatus{
		Status: *object,
	}
	payload, err := json.Marshal(withStatus)
	if err != nil {
		return "", err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return string(payload), nil
}
