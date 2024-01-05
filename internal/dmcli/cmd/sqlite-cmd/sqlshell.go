package sqlitecmd

import (
	"database/sql"
	"os"
	"reflect"
	"strings"
	"sync/atomic"
	"syscall"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

var (
	cmd         *cobra.Command
	db          *sql.DB
	isMultiLine atomic.Bool
	sqlBuilder  strings.Builder
	suggests    []prompt.Suggest
)

func init() {
	suggests = make([]prompt.Suggest, 0, len(sqlKeywords))
	for _, word := range sqlKeywords {
		suggests = append(suggests, prompt.Suggest{Text: word})
	}
}

func sqlShell(c *cobra.Command, sqliteDB *sql.DB) {
	cmd = c
	db = sqliteDB

	pt := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("> "),
		prompt.OptionPrefixTextColor(prompt.DefaultColor),
		prompt.OptionSwitchKeyBindMode(prompt.EmacsKeyBind),
		prompt.OptionCompletionOnDown(),
		// prompt.OptionShowCompletionAtStart(),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(buf *prompt.Buffer) {
				if err := syscall.Kill(os.Getpid(), syscall.SIGINT); err != nil {
					os.Exit(1)
				}
			},
		}),
		prompt.OptionLivePrefix(multiLinePrefix),
		// prompt.OptionOnExitCallback(),
	)
	pt.Run()
}

func executor(in string) {
	in = strings.Trim(in, " ")
	if len(in) == 0 {
		return
	}

	if sqlBuilder.Len() > 0 {
		sqlBuilder.WriteByte(' ')
	}
	sqlBuilder.WriteString(in)

	if !strings.HasSuffix(in, ";") {
		isMultiLine.Store(true)
		return
	}
	isMultiLine.Store(false)

	sqlStr := sqlBuilder.String()
	sqlBuilder.Reset()

	for _, s := range splitSql(sqlStr) {
		cmd.Println("sql: " + s)
		cmd.Println("  --------  data  --------  ")
		query(s)
		cmd.Println("")
	}
}

func query(sqlStr string) {
	rows, err := db.Query(sqlStr)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		for i, colName := range columns {
			var v interface{} = values[i]
			val := reflect.ValueOf(v)
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}
			cmd.Printf("%s: %v\n", colName, val.Interface())
		}
	}

	if err := rows.Err(); err != nil {
		cmd.PrintErrln(err)
		return
	}
}

func multiLinePrefix() (prefix string, useLivePrefix bool) {
	if isMultiLine.Load() {
		return " » ", true
	}

	return "", false
}

func completer(d prompt.Document) []prompt.Suggest {
	s := suggests
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
