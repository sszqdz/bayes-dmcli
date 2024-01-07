// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"

	"github.com/sszqdz/bayes-dmcli/internal/dmcli/cmd"
	"github.com/sszqdz/bayes-dmcli/internal/dmcli/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sszqdz/bayes-toolkit/environment"
)

const cmdName = "dmcli"

func loadConfig() *config.Config {
	var conf config.Config
	vtoml := viper.New()
	vtoml.SetConfigName(".dmconfig")
	vtoml.SetConfigType("toml")
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	if environment.LoadEnv().Is(environment.EnvDebug) {
		vtoml.AddConfigPath(filepath.Join(homeDir, "Documents", "projects", "bayes-dmcli", "config", cmdName))
	} else {
		workPath, err := os.Getwd()
		cobra.CheckErr(err)
		vtoml.AddConfigPath(filepath.Join(workPath))
		vtoml.AddConfigPath(filepath.Join(homeDir))
	}

	cobra.CheckErr(vtoml.ReadInConfig())
	cobra.CheckErr(vtoml.Unmarshal(&conf))

	return &conf
}

func main() {
	cobra.CheckErr(cmd.Execute(loadConfig()))
}
