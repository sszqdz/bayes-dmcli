// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"
	"path/filepath"

	"github.com/sszqdz/bayes-dmcli/internal/config"
	"github.com/sszqdz/bayes-dmcli/internal/root"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sszqdz/bayes-toolkit/environment"
)

func loadConfig() *config.Config {
	var conf config.Config
	vtoml := viper.New()
	vtoml.SetConfigName(".dmconfig")
	vtoml.SetConfigType("toml")
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	if environment.LoadEnv().Is(environment.EnvDebug) {
		vtoml.AddConfigPath(filepath.Join(homeDir, "Documents", "projects", "bayes-dmcli", "config"))
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

func Run() {
	cobra.CheckErr(root.Execute(loadConfig()))
}
