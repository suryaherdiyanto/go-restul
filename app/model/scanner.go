package model

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/go-restful/helper"
)

func ScanRow(d interface{}, rows *sql.Rows) {
	columns, err := rows.Columns()

	if err != nil {
		helper.ErrorPanic(err)
	}

	ref := reflect.TypeOf(d)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	if ref.Kind() != reflect.Struct {
		panic(fmt.Errorf("model.ScanRow only accepts struct, kind: %v", ref.Kind()))
	}

	var dRefs []interface{}

	for i := 0; i < ref.NumField(); i++ {
		for _, column := range columns {
			f := ref.Field(i)

			if f.Tag.Get("db") == column {
				dRefs = append(dRefs, reflect.ValueOf(d).Elem().Field(i).Addr().Interface())
			}
		}
	}

	rows.Scan(dRefs...)
}
