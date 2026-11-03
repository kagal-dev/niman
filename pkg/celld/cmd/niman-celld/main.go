// Command niman-celld is the sidecar niman asks about a cluster's celld
// nodes, running beside celld in the cluster's app.
package main

import (
	"os"

	"github.com/spf13/cobra"
)

// CmdName is the name of this executable.
const CmdName = "niman-celld"

var rootCmd = &cobra.Command{
	Use:          CmdName,
	Short:        "Tell niman about a cluster's celld nodes",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
}

func main() {
	// cobra prints the error itself; SilenceUsage keeps it from dumping
	// usage on failure.
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
