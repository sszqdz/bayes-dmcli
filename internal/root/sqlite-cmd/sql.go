// Copyright 2024 Moran. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// sqlKeywords REFERENCE LICENSE:
// https://github.com/dbcli/mycli/blob/main/mycli/sqlcompleter.py
// All rights reserved.

// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:

// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.

// * Redistributions in binary form must reproduce the above copyright notice, this
//   list of conditions and the following disclaimer in the documentation and/or
//   other materials provided with the distribution.

// * Neither the name of the {organization} nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// -------------------------------------------------------------------------------

// This program also bundles with it python-tabulate
// (https://pypi.python.org/pypi/tabulate) library. This library is licensed under
// MIT License.

// -------------------------------------------------------------------------------

package sqlitecmd

import (
	"slices"
	"strings"
)

var (
	sqlKeywords = []string{
		"SELECT", "FROM", "WHERE", "UPDATE", "DELETE FROM", "GROUP BY",
		"JOIN", "INSERT INTO", "LIKE", "LIMIT" /*"ACCESS",*/, "ADD", "ALL",
		"ALTER TABLE", "AND", "ANY", "AS", "ASC", "AUTO_INCREMENT",
		"BEFORE", "BEGIN", "BETWEEN", "BIGINT", "BINARY", "BY", "CASE",
		/*"CHANGE MASTER TO",*/ "CHAR", "CHARACTER SET", "CHECK", "COLLATE",
		"COLUMN", "COMMENT", "COMMIT", "CONSTRAINT", "CREATE", "CURRENT",
		"CURRENT_TIMESTAMP", "DATABASE", "DATE", "DECIMAL", "DEFAULT",
		"DESC", "DESCRIBE", "DROP", "ELSE", "END", "ENGINE", "ESCAPE",
		"EXISTS" /*"FILE",*/, "FLOAT", "FOR", "FOREIGN KEY", "FORMAT", "FULL",
		"FUNCTION", "GRANT", "HAVING", "HOST", "IDENTIFIED", "IN",
		"INCREMENT", "INDEX", "INT", "INTEGER", "INTERVAL", "INTO", "IS",
		"KEY", "LEFT", "LEVEL", "LOCK", "LOGS", "LONG", /*"MASTER",*/
		"MEDIUMINT", "MODE", "MODIFY", "NOT", "NULL", "NUMBER", "OFFSET",
		"ON", "OPTION", "OR", "ORDER BY", "OUTER", "OWNER", "PASSWORD",
		"PORT", "PRIMARY", "PRIVILEGES", "PROCESSLIST", "PURGE",
		"REFERENCES", "REGEXP", "RENAME", "REPAIR", "RESET", "REVOKE",
		"RIGHT", "ROLLBACK", "ROW", "ROWS", "ROW_FORMAT", "SAVEPOINT",
		"SESSION", "SET", "SHARE", "SHOW", "SLAVE", "SMALLINT", "SMALLINT",
		"START", "STOP", "TABLE", "THEN", "TINYINT", "TO", "TRANSACTION",
		"TRIGGER", "TRUNCATE", "UNION", "UNIQUE", "UNSIGNED", "USE",
		"USER", "USING", "VALUES", "VARCHAR", "VIEW", "WHEN", "WITH",
	}
	// TODO In this case, we use "main.xxx" instead of just "xxx" to construct the prompt.
	// The reason for this is that when the github.com/c-bata/go-prompt selects a certain prompt,
	// it splits the text by spaces and then performs an overwrite.
	// If the prompt matching is not done based on space-separated words,
	// it can result in the prompt overwriting more than intended.
	sqliteKeywords = []string{
		"SQLite", "PRAGMA", "ATTACH", "DETACH",
		"schema", "sqlite_version", "main",
		"main.table_list", "main.table_info", "main.table_xinfo",
		"main.index_list", "main.index_info", "main.index_xinfo",
	}
	queryKeywords = []string{"SELECT", "PRAGMA"}
)

type sqlInfo struct {
	SQL     string
	IsQuery bool
}

func splitSql(s string) []*sqlInfo {
	strs := strings.SplitAfter(s, ";")
	ret := make([]*sqlInfo, 0, len(strs))
	for _, str := range strs {
		str = strings.TrimSpace(str)
		if str == "" || str == ";" {
			continue
		}
		before, after, found := strings.Cut(str, " ")
		if found {
			before = strings.ToUpper(before)
		}
		ret = append(ret, &sqlInfo{
			SQL:     before + " " + after,
			IsQuery: slices.Contains[[]string](queryKeywords, before),
		})
	}

	return ret
}
