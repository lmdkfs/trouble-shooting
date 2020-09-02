package cmd

import (
	"errors"
	"fmt"
	"trouble-shooting/cmd/api"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "gin-admin",
	Short:        "gin-admin",
	SilenceUsage: true,
	Long:         `gin-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tips()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tips()
	},
}

func tips() {
	usageStr := `欢迎使用` + `gin-admin` + `可以使用` + `-h`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
