package vo

import (
	"encoding/json"

	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func TranslateToDockerHost(payload []byte) (docker.DockerHost, error) {
	blankObject := docker.DockerHost{}
	resultObject := docker.DockerHost{}
	err := json.Unmarshal(payload, &resultObject)
	//client.Logger.Info(fmt.Sprintf("Watcher   #%v ", watcher))
	if err != nil {
		//client.Logger.Info(fmt.Sprintf("Error loading CRD   #%v ", err))
		return blankObject, err
	}
	//client.Logger.Info(fmt.Sprintf("Update status:  %#v ", client.CRD))
	return resultObject, nil
}
