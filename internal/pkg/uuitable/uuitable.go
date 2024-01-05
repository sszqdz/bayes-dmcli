package uuitable

import (
	"fmt"
	"strings"

	"github.com/gosuri/uitable"
)

func AddHeader(tb *uitable.Table, data ...interface{}) {
	tb.AddRow(data...)
	var separators []interface{}
	for _, d := range data {
		length := len(fmt.Sprintf("%v", d))
		separators = append(separators, strings.Repeat("-", length))
	}
	tb.AddRow(separators...)
}
