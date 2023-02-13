package vo

import (
	"encoding/json"
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

type WithStatus struct {
	Status docker.DockerHostStatus `json:"status,omitempty"`
}

func TranslateToDockerComposeRunner(payload []byte) (docker.DockerComposeRunner, error) {
	logutils.Logger.Info(string(payload))
	blankObject := docker.DockerComposeRunner{}
	resultObject := docker.DockerComposeRunner{}
	err := json.Unmarshal(payload, &resultObject)
	logutils.Logger.Info(fmt.Sprintf("%#v", resultObject))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
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
		logutils.Logger.Error(err.Error())
		return "", err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return string(payload), nil
}

func TranslateToComposeStatusArray(payload []byte) ([]docker.ComposeStatus, error) {
	blankObject := []docker.ComposeStatus{}
	resultObject := []docker.ComposeStatus{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}
