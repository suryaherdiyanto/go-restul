package model

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-restful/helper"
)

func ScanStruct(d interface{}, rows *sql.Rows) error {
	columns, err := rows.Columns()

	if err != nil {
		helper.ErrorPanic(err)
	}

	ref := reflect.TypeOf(d)
	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	refKind := ref.Kind()
	if refKind != reflect.Struct {
		return errors.New(fmt.Sprintf("model.ScanRow only accepts struct kind: %v", refKind))
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

	return nil

}

func ScanMap(d interface{}, rows *sql.Rows) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	ref := reflect.TypeOf(d)

	if ref.Kind() == reflect.Ptr {
		ref = ref.Elem()
	}

	refKind := ref.Kind()

	if refKind != reflect.Map {
		return errors.New(fmt.Sprintf("model.ScanMap only accepts map kind: %v", refKind))
	}

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
		case reflect.Bool:
			reflect.ValueOf(d).Elem().SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(value.Interface().(bool)))
		default:
			reflect.ValueOf(d).Elem().SetMapIndex(reflect.ValueOf(column), reflect.ValueOf(string(value.Interface().([]byte))))
		}
	}

	return nil
}
