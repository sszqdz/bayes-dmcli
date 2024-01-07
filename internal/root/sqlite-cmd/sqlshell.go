// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package sqlitecmd

import (
	"database/sql"
	"os"
	"strings"
	"sync/atomic"
	"syscall"

	"github.com/sszqdz/bayes-dmcli/internal/pkg/uuitable"

	"github.com/c-bata/go-prompt"
	"github.com/gosuri/uitable"
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
	suggests = make([]prompt.Suggest, 0, len(sqlKeywords)+len(sqliteKeywords))
	for _, word := range sqlKeywords {
		suggests = append(suggests, prompt.Suggest{Text: word})
	}
	for _, word := range sqliteKeywords {
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
		prompt.OptionMaxSuggestion(10),
		prompt.OptionCompletionOnDown(),
		// prompt.OptionShowCompletionAtStart(),
		prompt.OptionLivePrefix(multiLinePrefix),
		prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(buf *prompt.Buffer) {
				if err := syscall.Kill(os.Getpid(), syscall.SIGINT); err != nil {
					os.Exit(1)
				}
			},
		}),
		prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
			if breakline {
				return false
			}
			in = strings.TrimSpace(in)
			if !strings.HasSuffix(in, ";") {
				return false
			}
			return strings.TrimSpace(sqlBuilder.String())+in == "exit;"
		}),
		// prompt.OptionOnExitCallback(),
	)

	pt.Run()
}

func executor(in string) {
	in = strings.TrimSpace(in)
	if len(in) == 0 {
		return
	}
	if sqlBuilder.Len() > 0 {
		sqlBuilder.WriteByte(' ')
	}
	// determine multi-line
	if !strings.HasSuffix(in, ";") {
		if beforeStr, backslashFound := strings.CutSuffix(in, `\`); backslashFound {
			in = beforeStr
		}
		sqlBuilder.WriteString(in)
		isMultiLine.Store(true)
		return
	}
	// end line
	sqlBuilder.WriteString(in)
	isMultiLine.Store(false)

	sqlStr := sqlBuilder.String()
	sqlBuilder.Reset()
	if len(sqlStr) == 0 {
		return
	}

	for _, si := range splitSql(sqlStr) {
		if si.IsQuery {
			cmd.Println("QUERY :" + si.SQL)
			cmd.Println("________DATA________")
			query(si.SQL)
		} else {
			cmd.Println("EXEC: " + si.SQL)
			cmd.Println("________RESULT________")
			exec(si.SQL)
		}
		cmd.Println()
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
	if len(columns) == 0 {
		return
	}

	columnSlice := make([]any, len(columns))
	for i, column := range columns {
		columnSlice[i] = column
	}
	tb := newTable(50)
	uuitable.AddHeader(tb, columnSlice...)

	hasData := false
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		hasData = true
		data := make([]interface{}, 0, len(columns))
		for i := range columns {
			var v interface{} = values[i]
			data = append(data, v)
		}
		tb.AddRow(data...)
	}
	cmd.Println(tb)
	if !hasData {
		cmd.Println("__NO__DATA__")
	}

	if err := rows.Err(); err != nil {
		cmd.PrintErrln(err)
		return
	}
}

func exec(sqlStr string) {
	result, err := db.Exec(sqlStr)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		cmd.PrintErrln(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		cmd.PrintErrln(err)
	}
	tb := newTable(50)
	uuitable.AddHeader(tb, "lastInsertId", "rowsAffected")
	tb.AddRow(lastInsertId, rowsAffected)
	cmd.Println(tb)
}

func multiLinePrefix() (prefix string, useLivePrefix bool) {
	if isMultiLine.Load() {
		return " » ", true
	}

	return "", false
}

func completer(d prompt.Document) []prompt.Suggest {
	return prompt.FilterHasPrefix(suggests, d.GetWordBeforeCursor(), true)
}

func newTable(maxColWidth uint, wrap ...bool) *uitable.Table {
	tb := uitable.New()
	tb.Separator = "  "
	tb.MaxColWidth = maxColWidth
	tb.Wrap = true
	if len(wrap) == 1 {
		tb.Wrap = wrap[0]
	}
	return tb
}
