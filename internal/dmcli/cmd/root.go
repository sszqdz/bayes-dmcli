package cmd

import (
	sqlitecmd "bayes/dmcli/internal/dmcli/cmd/sqlite-cmd"
	"bayes/dmcli/internal/dmcli/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dmcli",
	Short: "Database Command Tool",
	Long:  `A database command line tool, supports sqlite, redis and more!`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "v0.1.0",
}

func Execute(conf *config.Config) error {
	rootCmd.AddCommand(sqlitecmd.Commands(conf)...)
	return rootCmd.Execute()
}