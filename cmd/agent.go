/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/6zacode-toolbox/docker-agent/internal/controller"
	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	"github.com/spf13/cobra"
)

var crdAPIVersion string
var crdNamespace string
var crdName string
var crdResource string

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Execute test Agent",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		crdConfig := &crdtools.CRDConfig{
			APIVersion:   crdAPIVersion,
			Namespace:    crdNamespace,
			InstanceName: crdName,
			Resource:     crdResource,
		}
		logutils.Logger.Info("Execute Agent")
		controller.ExecuteAgent(crdConfig)

	},
}

func init() {
	agentCmd.Flags().StringVar(&crdAPIVersion, "crd-api-version", "tool.6zacode-toolbox.github.io/v1", `CRD API Version.`)
	agentCmd.Flags().StringVar(&crdNamespace, "crd-namespace", "default", `CRD Namespace.`)
	agentCmd.Flags().StringVar(&crdName, "crd-instance", "", `CRD instance name. Mandatory in CRD mode.`)
	agentCmd.Flags().StringVar(&crdResource, "crd-resource", "dockerhosts", `CRD Resource name`)
	rootCmd.AddCommand(agentCmd)
}
