package cmd

import(
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "drift",
    Short: "short",
    Long: `long`,
}
