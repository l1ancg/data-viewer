package utils

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

func Parse(sql string, param map[string]interface{}) (string, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		fmt.Println("Parse error: ", err.Error())
		return "", err
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		// for _, expr := range stmt.SelectExprs {
		// 	switch expr := expr.(type) {
		// 	case *sqlparser.AliasedExpr:
		// 		fmt.Println(expr.Expr)
		// 	}
		// }
		stmt.Where = processWhere(stmt.Where, param)
	default:
		panic("unknown sql type")
	}
	return sqlparser.String(stmt), nil
}

func processWhere(where *sqlparser.Where, param map[string]interface{}) *sqlparser.Where {
	if where == nil {
		return nil
	}

	switch expr := where.Expr.(type) {
	case *sqlparser.AndExpr:
		expr.Left = processWhere(&sqlparser.Where{Expr: expr.Left}, param).Expr
		expr.Right = processWhere(&sqlparser.Where{Expr: expr.Right}, param).Expr
	case *sqlparser.ComparisonExpr:
		col := sqlparser.String(expr.Left)
		if val, ok := param[col]; ok {
			switch v := val.(type) {
			case string:
				expr.Right = sqlparser.NewStrVal([]byte(v))
			case int:
				expr.Right = sqlparser.NewIntVal([]byte(fmt.Sprintf("%d", v)))
			default:
				return nil
			}
		} else {
			return nil
		}
	default:
		// handle other cases if needed
	}

	return where
}
