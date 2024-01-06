package main

import (
	"bayes-dmcli/internal/dmcli/cmd"
	"bayes-dmcli/internal/dmcli/config"

	"os"
	"path/filepath"

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
	confPath := filepath.Join(homeDir)
	if environment.LoadEnv().Is(environment.EnvDebug) {
		confPath = filepath.Join(homeDir, "Documents", "projects", "bayes-dmcli", "config", cmdName)
	}
	vtoml.AddConfigPath(confPath)
	cobra.CheckErr(vtoml.ReadInConfig())
	cobra.CheckErr(vtoml.Unmarshal(&conf))

	return &conf
}

func main() {
	cobra.CheckErr(cmd.Execute(loadConfig()))
}
