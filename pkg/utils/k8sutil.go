package utils

import (
    cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func GetFactory() *cmdutil.Factory {
	factory := cmdutil.NewFactory(nil)
	return factory
}
