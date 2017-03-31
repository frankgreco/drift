package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/frankgreco/drift/etcd"
)

var startCmd = &cobra.Command{
    Use:   `start`,
    Short: `short`,
    Long:  `long`,
    Run: func(cmd *cobra.Command, args []string) {
        master, err := cmd.Flags().GetString("master")
		cfg := controller.NewControllerConfig(master, "")
        fmt.Println(cfg.Namespace)




        etcd.Watch()
    },
}

func init() {
    startCmd.Flags().BoolP("master", "", "", "The address of the Kubernetes API server (overrides any value in kubeconfig)")
    startCmd.Flags().BoolP("etcd-servers", "", "http://127.0.0.1:2379", "List of etcd servers to connect with (scheme://ip:port), comma separated.")
}
