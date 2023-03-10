package controller

import (
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/dockerutils"
	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/internal/vo"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	"github.com/6zacode-toolbox/docker-agent/pkg/k8s"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func ExecuteAgent(crd *crdtools.CRDConfig) error {
	crdContent, _ := crdtools.GetCRD(crd)
	dockerHost, _ := vo.TranslateToDockerHost(crdContent)
	logutils.Logger.Info(fmt.Sprintf("%#v", dockerHost))

	dockerInfo, _ := dockerutils.ExecuteDockerInfo()
	dockerPS, _ := dockerutils.ExecuteDockerPS()

	status := &docker.DockerHostStatus{
		Containers: dockerPS,
		DockerHost: dockerInfo,
		Instanced:  true,
	}

	payload, err := vo.FromDockerHostStatus(status)
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	crdtools.UpdateStatus(crd, payload)
	return nil
}

func ExecuteDockerComposeRunnerUp(crd *crdtools.CRDConfig) error {
	crdContent, err := crdtools.GetCRD(crd)
	logutils.Logger.Info(string(crdContent))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	composeRunner, err := vo.TranslateToDockerComposeRunner(crdContent)
	logutils.Logger.Info(fmt.Sprintf("%#v", composeRunner))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	logutils.Logger.Info(fmt.Sprintf("%#v", composeRunner))

	// Do something
	// Execute Logic
	dockerRunnerStatus, err := dockerutils.ExecuteCompose()
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	payload, err := vo.FromDockerComposeRunnerStatus(&dockerRunnerStatus)
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	crdtools.UpdateStatus(crd, payload)
	return nil
}

func ExecuteDockerComposeRunnerDown(crd *crdtools.CRDConfig) error {
	// Execute Logic
	_, err := dockerutils.ExecuteCompose()
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	err = k8s.DeleteConfigMap(crd)
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	return nil
}
