package k8s

import (
	"context"
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteConfigMap(config *crdtools.CRDConfig) error {
	clientset, err := crdtools.GetK8SConfig()
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}

	err = clientset.CoreV1().ConfigMaps(config.Namespace).DeleteCollection(context.TODO(), metaV1.DeleteOptions{}, metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("instance=%s,resouce-owner=%s", config.InstanceName, config.Resource),
	})
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	return nil
}
