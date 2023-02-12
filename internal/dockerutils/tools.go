package dockerutils

import (
	"fmt"
	"os/exec"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/internal/vo"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func ExecuteDockerInfo() (docker.DockerInfo, error) {
	blankObject := docker.DockerInfo{}

	cmd := exec.Command("/home/app/docker_info.sh")
	stdout, err := cmd.Output()
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	result, err := vo.TranslateToDockerInfo(stdout)
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	return result, nil
}

func ExecuteDockerPS() ([]docker.DockerContainerSummary, error) {
	blankObject := []docker.DockerContainerSummary{}

	cmd := exec.Command("/home/app/docker_ps.sh")
	stdout, err := cmd.Output()
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	result, err := vo.TranslateToDockerContainerSummarys(stdout)
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	return result, nil
}

func ExecuteCompose(runner *docker.DockerComposeRunner) (docker.DockerComposeRunnerStatus, error) {
	blankObject := docker.DockerComposeRunnerStatus{}
	// Setup variables for SHELL that are not already on the pod by the controller
	// Controller should provide such variables
	cmd := exec.Command("/home/app/docker_compose.sh")
	stdout, err := cmd.Output()
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	logutils.Logger.Info(fmt.Sprintf("%#v", stdout))
	//result := docker.DockerComposeRunnerStatus{}
	/*
		result, err := vo.TranslateToDockerContainerSummarys(stdout)
		if err != nil {
			logutils.Logger.Error(fmt.Sprintf("%#v",err))
			return blankObject, err
		}
	*/
	return blankObject, nil

}
