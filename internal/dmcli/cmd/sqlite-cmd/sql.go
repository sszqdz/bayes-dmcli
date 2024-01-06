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
