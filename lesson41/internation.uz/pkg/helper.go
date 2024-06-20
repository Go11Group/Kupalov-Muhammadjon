package pkg

import "fmt"

// formatting queries 
func AppendParamPlaceholder(query *string, paramCount int) {
	if paramCount > 0 {
		*query += fmt.Sprintf("$%d", paramCount)
	}
}