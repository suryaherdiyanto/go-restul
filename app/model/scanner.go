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

	refKind := ref.Kind()
	if refKind != reflect.Struct && refKind != reflect.Map {
		panic(fmt.Errorf("model.ScanRow only accepts struct, or map kind: %v", refKind))
	}

	if refKind == reflect.Map {
		values := make([]interface{}, len(columns))
		for i := range values {
			var v interface{}
			values[i] = &v
		}
		if err := rows.Scan(values...); err != nil {
			panic(err)
		}

		for i, column := range columns {
			value := reflect.ValueOf(values[i]).Elem()
			switch value.Elem().Kind() {
			case reflect.Int64:
				reflect.ValueOf(d).Elem().SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(int(value.Interface().(int64))))
			case reflect.Float64:
				reflect.ValueOf(d).Elem().SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(float32(value.Interface().(float64))))
			default:
				reflect.ValueOf(d).Elem().SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(string(value.Interface().([]byte))))
			}
		}
	}

	if refKind == reflect.Struct {
		var dRefs []interface{}
		for _, column := range columns {
			for i := 0; i < ref.NumField(); i++ {
				f := ref.Field(i)

				if f.Tag.Get("db") == column {
					dRefs = append(dRefs, reflect.ValueOf(d).Elem().Field(i).Addr().Interface())
				}
			}
		}
		rows.Scan(dRefs...)
	}

}
