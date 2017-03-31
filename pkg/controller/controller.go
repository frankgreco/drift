package controller

import (
    "fmt"

    "github.com/frankgreco/drift/pkg/utils"
    "k8s.io/kubernetes/pkg/client/unversioned"
)

const (
	tprName = "drift.k8s.io"
)

type Config struct {
	Namespace  string
	KubeCli    *unversioned.Client
	MasterHost string
}

func NewControllerConfig(masterHost, ns string) Config {
	f := utils.GetFactory()
	kubecli, err := f.Client()
	if err != nil {
		fmt.Errorf("Can not get kubernetes config: %s", err)
	}
	if ns == "" {
		ns, _, err = f.DefaultNamespace()
		if err != nil {
			fmt.Errorf("Can not get kubernetes config: %s", err)
		}
	}
	if masterHost == "" {
		k8sConfig, err := f.ClientConfig()
		if err != nil {
			fmt.Errorf("Can not get kubernetes config: %s", err)
		}
		if k8sConfig == nil {
			fmt.Errorf("Got nil k8sConfig, please check if k8s cluster is available.")
		} else {
			masterHost = k8sConfig.Host
		}
	}

	cfg := Config{
		Namespace:  ns,
		KubeCli:    kubecli,
		MasterHost: masterHost,
	}

	return cfg
}
