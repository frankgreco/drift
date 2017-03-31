package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use:   `version`,
    Short: `short`,
    Long:  `long`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("v1.0.0")
    },
}
