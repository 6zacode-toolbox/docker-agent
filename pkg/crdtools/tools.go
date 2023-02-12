package crdtools

import (
	"context"
	"fmt"

	"github.com/6zacode-toolbox/docker-agent/internal/logutils"
	"go.uber.org/zap"
	api "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var Logger *zap.Logger

//:= logutils.InitializeLogger()

type CRDConfig struct {
	APIVersion   string
	Namespace    string
	InstanceName string
	Resource     string
}

func GetK8SConfig() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	logutils.Logger.Info("Config used")
	return clientSet, nil

}

func UpdateStatus(CRD *CRDConfig, statusPayload string) error {
	//Logger.Debug(fmt.Sprintf("Watcher Config: #%v ", CRD))
	clientSet, _ := GetK8SConfig()
	//myPatch := fmt.Sprintf(`{"status":{"status":"%s"}}`, statusPayload)
	logutils.Logger.Info(fmt.Sprintf("%#v", statusPayload))
	//Logger.Debug(fmt.Sprintf("Watcher Patch: #%v ", myPatch))
	_, err := clientSet.RESTClient().
		Patch(api.MergePatchType).
		AbsPath("/apis/" + CRD.APIVersion).
		SubResource("status").
		Namespace(CRD.Namespace).
		Resource(CRD.Resource).
		Name(CRD.InstanceName).
		Body([]byte(statusPayload)).
		DoRaw(context.TODO())
	if err != nil {
		//Logger.Info(fmt.Sprintf("Error updating CRD   #%v ", err))
		return err
	}
	//Logger.Info(fmt.Sprintf("Update status:  %#v ", CRD))
	return nil
}

// GetCRD returns the Watcher CRD Object, used to feed configs
func GetCRD(CRD *CRDConfig) ([]byte, error) {
	//Logger.Debug(fmt.Sprintf("CRD: #%v ", CRD))
	clientSet, _ := GetK8SConfig()
	object, _ := clientSet.RESTClient().
		Get().
		AbsPath("/apis/" + CRD.APIVersion).
		Namespace(CRD.Namespace).
		Resource(CRD.Resource).
		Name(CRD.InstanceName).
		DoRaw(context.TODO())
	//Logger.Info(fmt.Sprintf("Get CRD   #%v ", string(object)))
	return object, nil

}
