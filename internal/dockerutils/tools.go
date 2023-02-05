package dockerutils

import (
	"fmt"
	"os/exec"

	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
)

func ExecuteDockerInfo() (docker.DockerInfo, error) {
	blankObject := docker.DockerInfo{}

	cmd := exec.Command("/home/app/docker_info.sh")
	stdout, err := cmd.Output()
	fmt.Println(string(stdout))
	if err != nil {
		fmt.Println(err.Error())
		return blankObject, err
	}
	fmt.Print(string(stdout))
	return blankObject, nil
}
