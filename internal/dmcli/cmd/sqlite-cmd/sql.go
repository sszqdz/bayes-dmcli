package sqlitecmd

import (
	"strings"
)

var sqlKeywords = []string{
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

func splitSql(s string) []string {
	strs := strings.SplitAfter(s, ";")
	ret := make([]string, 0, len(strs))
	for _, str := range strs {
		str = strings.Trim(str, " ")
		if str == "" || str == ";" {
			continue
		}
		ret = append(ret, str)
	}

	return ret
}
