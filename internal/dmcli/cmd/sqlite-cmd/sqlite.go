package sqlitecmd

import (
	"bayes/dmcli/internal/dmcli/config"

	"github.com/spf13/cobra"
)

var conf *config.Config

func Commands(c *config.Config) []*cobra.Command {
	conf = c
	return []*cobra.Command{sqliteCmd}
}

var sqliteCmd = &cobra.Command{
	Use:   "sqlite",
	Short: "Sqlite Command Tool",
	Long:  `A sqlite command line tool, supports sqlite, redis and more!`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	if conf != nil {
		cmd.Printf("test...%v\n", conf.DatabaseList)
	}
}
