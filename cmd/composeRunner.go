/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/6zacode-toolbox/docker-agent/internal/controller"
	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	docker "github.com/6zacode-toolbox/docker-operator/operator/api/v1"
	"github.com/spf13/cobra"
)

// composeRunnerCmd represents the composeRunner command
var composeRunnerCmd = &cobra.Command{
	Use:   "compose-runner",
	Short: "",
	Long:  ` `,
	Run: func(cmd *cobra.Command, args []string) {
		logutils.Logger.Info("composeRunner called")
		crdConfig := &crdtools.CRDConfig{
			APIVersion:   crdAPIVersion,
			Namespace:    crdNamespace,
			InstanceName: crdName,
			Resource:     crdResource,
		}
		logutils.Logger.Info("Execute Agent")
		logutils.Logger.Info(crdConfig.APIVersion)
		if os.Getenv("ACTION") == docker.COMPOSE_ACTION_DOWN {
			controller.ExecuteDockerComposeRunnerDown(crdConfig)
		} else {
			controller.ExecuteDockerComposeRunnerUp(crdConfig)
		}
	},
}

func init() {

	composeRunnerCmd.Flags().StringVar(&crdAPIVersion, "crd-api-version", "tool.6zacode-toolbox.github.io/v1", `CRD API Version.`)
	composeRunnerCmd.Flags().StringVar(&crdNamespace, "crd-namespace", "default", `CRD Namespace.`)
	composeRunnerCmd.Flags().StringVar(&crdName, "crd-instance", "", `CRD instance name. Mandatory in CRD mode.`)
	composeRunnerCmd.Flags().StringVar(&crdResource, "crd-resource", "dockercomposerunners", `CRD Resource name`)

	rootCmd.AddCommand(composeRunnerCmd)
}
