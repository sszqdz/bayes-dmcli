package sqlitecmd

import (
	"bayes-dmcli/internal/dmcli/config"
	"bayes-dmcli/internal/pkg/uuitable"
	"database/sql"
	"errors"

	"github.com/c-bata/go-prompt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	_ "github.com/glebarez/go-sqlite"
)

var (
	conf      *config.Config
	sqliteCmd = &cobra.Command{
		Use:   "sqlite",
		Short: "Sqlite Command Tool",
		Long:  `A sqlite command line tool, supports sqlite, redis and more!`,
		Run:   sqliteCmdRun,
	}
)

func Commands(c *config.Config) []*cobra.Command {
	conf = c
	return []*cobra.Command{sqliteCmd}
}

func sqliteCmdRun(cmd *cobra.Command, args []string) {
	// select db
	dbConf, err := selectDB(cmd)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	// open db
	db, err := openDB(cmd, dbConf)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	defer db.Close()
	cmd.Println("sqlite connected!")
	// sql shell
	sqlShell(cmd, db)
}

func selectDB(cmd *cobra.Command) (*config.Database, error) {
	tb := uitable.New()
	tb.MaxColWidth = 20
	tb.Wrap = true

	uuitable.AddHeader(tb, "#", "Name", "Desc", "Driver", "Source", "MaxIdleConn", "MaxOpenConn", "ConnMaxLifetime")
	i := 1
	for _, dbConf := range conf.DatabaseList {
		if dbConf.Driver != "sqlite" {
			continue
		}
		tb.AddRow(i, dbConf.Name, dbConf.Desc, dbConf.Driver, dbConf.Source, dbConf.MaxIdleConn, dbConf.MaxOpenConn, dbConf.ConnMaxLifetime)
		i++
	}
	cmd.Println(tb)
	// waiting input
	indexStr := prompt.Input("> ", func(prompt.Document) []prompt.Suggest { return nil })
	index, err := cast.ToIntE(indexStr)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	i = 1
	for _, dbConf := range conf.DatabaseList {
		if dbConf.Driver != "sqlite" {
			continue
		}
		if i == index {
			cmd.Println("You select ", i)
			return dbConf, nil
		}
		i++
	}

	return nil, errors.New("not exists")
}

func openDB(cmd *cobra.Command, dbConf *config.Database) (*sql.DB, error) {
	sqlDB, err := sql.Open(dbConf.Driver, dbConf.Source)
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConn)
	sqlDB.SetConnMaxIdleTime(0)
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(dbConf.ConnMaxLifetime)

	return sqlDB, nil
}
