package main

import (
	"bayes-dmcli/internal/dmcli/cmd"
	"bayes-dmcli/internal/dmcli/config"

	"fmt"
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
	if environment.LoadEnv().Is(environment.EnvDebug) {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)
		vtoml.AddConfigPath(filepath.Join(homeDir, "Documents", "projects", "bayes-dmcli", "config", cmdName))
	} else {
		vtoml.AddConfigPath(filepath.Join("/app", cmdName, "config"))
	}
	cobra.CheckErr(vtoml.ReadInConfig())
	cobra.CheckErr(vtoml.Unmarshal(&conf))

	return &conf
}

func main() {
	conf := loadConfig()
	if conf != nil {
		for _, db := range conf.DatabaseList {
			fmt.Printf("driver: %s\n", db.Driver)
		}
		for _, rds := range conf.RedisList {
			fmt.Printf("rds: %s\n", rds.Addr)
		}
	}
	cobra.CheckErr(cmd.Execute(conf))
}
