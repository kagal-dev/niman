// Command niman manages clusters of Workers and Durable Objects, whether
// they run on celld or on Cloudflare.
package main

// cspell:words celld

import (
	"os"

	"github.com/spf13/cobra"
)

// CmdName is the name of this executable.
const CmdName = "niman"

var rootCmd = &cobra.Command{
	Use:          CmdName,
	Short:        "Manage clusters of Workers and Durable Objects",
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
