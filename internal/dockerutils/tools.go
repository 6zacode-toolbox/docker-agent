package dockerutils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/internal/vo"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
	funk "github.com/thoas/go-funk"
)

func ExecuteDockerInfo() (docker.DockerInfo, error) {
	blankObject := docker.DockerInfo{}

	cmd := exec.Command("/home/app/docker_info.sh")
	stdout, err := cmd.Output()
	logutils.Logger.Info(string(stdout))
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	logutils.Logger.Info(string(stdout))
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
	logutils.Logger.Info(string(stdout))
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	logutils.Logger.Info(string(stdout))
	result, err := vo.TranslateToDockerContainerSummarys(stdout)
	if err != nil {
		logutils.Logger.Error(fmt.Sprintf("%#v", err))
		return blankObject, err
	}
	return result, nil
}

func ExecuteCompose() (docker.DockerComposeRunnerStatus, error) {
	blankObject := docker.DockerComposeRunnerStatus{}
	// Setup variables for SHELL that are not already on the pod by the controller
	// Controller should provide such variables
	cmd := exec.Command("/home/app/docker_compose.sh")
	stdout, err := cmd.Output()
	logutils.Logger.Info(string(stdout))
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}
	logutils.Logger.Info(string(stdout))
	beforeStatus, err := ReadComposeStatus("/var/tmp/before.json")
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}

	afterStatus, err := ReadComposeStatus("/var/tmp/after.json")
	if err != nil {
		logutils.Logger.Error(err.Error())
		return blankObject, err
	}

	//mappingBefore := funk.ToMap(beforeStatus, "ConfigFiles")
	//keysBefore := reflect.ValueOf(mappingBefore).MapKeys()
	var oldStatus []string
	for _, v := range beforeStatus {
		oldStatus = append(oldStatus, v.Name)
	}

	var newStatus []docker.ComposeStatus
	logutils.Logger.Info(fmt.Sprintf("keysBefore: %#v", oldStatus))
	for _, v := range afterStatus {
		logutils.Logger.Info(fmt.Sprintf("v.ConfigFiles: %#v", v.Name))
		if !funk.Contains(oldStatus, v.Name) {
			newStatus = append(newStatus, v)
			logutils.Logger.Info(fmt.Sprintf("newStatus: %#v", v))
		}

	}

	result := docker.DockerComposeRunnerStatus{
		ComposeStatus: newStatus,
		Instanced:     true,
	}
	return result, nil

}

func ReadComposeStatus(file string) ([]docker.ComposeStatus, error) {
	afterComposeStatus, err := os.ReadFile(file)
	if err != nil {
		logutils.Logger.Error(err.Error())
		return nil, err
	}
	logutils.Logger.Info(string(afterComposeStatus))
	status, err := vo.TranslateToComposeStatusArray(afterComposeStatus)
	if err != nil {
		logutils.Logger.Error(err.Error())
		return nil, err
	}
	return status, nil
}
