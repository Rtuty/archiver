package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "You are welcome to simple archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		handleErr(err)
	}
}

func handleErr(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
