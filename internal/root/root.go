// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package root

import (
	"github.com/sszqdz/bayes-dmcli/internal/config"
	sqlitecmd "github.com/sszqdz/bayes-dmcli/internal/root/sqlite-cmd"

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
