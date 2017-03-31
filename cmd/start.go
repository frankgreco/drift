package cmd

import (
    "github.com/spf13/cobra"
    "github.com/frankgreco/drift/etcd"
)

func init() {
    RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
    Use:   `start`,
    Short: `short`,
    Long:  `long`,
    Run: func(cmd *cobra.Command, args []string) {
        etcd.Watch()
    },
}
