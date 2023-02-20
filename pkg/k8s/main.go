package k8s

import (
	"context"
	"fmt"
	"os"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"github.com/6zacode-toolbox/docker-agent/pkg/crdtools"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func DeleteConfigMap(config *crdtools.CRDConfig) error {
	namespace := os.Getenv("OPERATOR_NAMESPACE")
	if namespace == "" {
		namespace = config.Namespace
	}
	clientset, err := crdtools.GetK8SConfig()
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	logutils.Logger.Info(fmt.Sprintf("try to remove configmap: %s,%s", config.InstanceName, config.Resource))
	err = clientset.CoreV1().ConfigMaps(namespace).DeleteCollection(context.TODO(), metaV1.DeleteOptions{}, metaV1.ListOptions{
		LabelSelector: fmt.Sprintf("instance=%s,resouce-owner=%s", config.InstanceName, config.Resource),
	})
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	return nil
}

func DeleteConfigMapByName(name string, namespace string) error {
	clientset, err := crdtools.GetK8SConfig()
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	logutils.Logger.Info(fmt.Sprintf("try to remove configmap: %s,%s", name, namespace))
	err = clientset.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
	if err != nil {
		logutils.Logger.Error(err.Error())
		return err
	}
	logutils.Logger.Info(fmt.Sprintf("configmap removed: %s,%s", name, namespace))
	return nil
}
