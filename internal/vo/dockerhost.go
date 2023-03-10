package vo

import (
	"encoding/json"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func TranslateToDockerHost(payload []byte) (docker.DockerHost, error) {
	blankObject := docker.DockerHost{}
	resultObject := docker.DockerHost{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}

func TranslateToDockerInfo(payload []byte) (docker.DockerInfo, error) {
	blankObject := docker.DockerInfo{}
	resultObject := docker.DockerInfo{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}

func TranslateToDockerContainerSummarys(payload []byte) ([]docker.DockerContainerSummary, error) {
	blankObject := []docker.DockerContainerSummary{}
	resultObject := []docker.DockerContainerSummary{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}

func FromDockerHostStatus(object *docker.DockerHostStatus) (string, error) {
	type WithStatus struct {
		Status docker.DockerHostStatus `json:"status,omitempty"`
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
