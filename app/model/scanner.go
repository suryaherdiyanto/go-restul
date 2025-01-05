package model

import (
	"database/sql"
	"reflect"

	"github.com/go-restful/helper"
)

func ScanRow(d interface{}, rows *sql.Rows) {
	columns, err := rows.Columns()

	if err != nil {
		helper.ErrorPanic(err)
	}

	ref := reflect.TypeOf(d)

	for i := 0; i < ref.NumField(); i++ {
		for _, column := range columns {
			f := ref.Field(i)

			if f.Tag.Get("db") == column {
				rows.Scan(reflect.ValueOf(f).Addr())
			}
		}
	}
}
