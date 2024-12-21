package helper

import (
	"database/sql"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ErrorPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HandleNullString(val interface{}) sql.NullString {
	if val == nil {
		return sql.NullString{String: "", Valid: false}
	}

	return sql.NullString{String: val.(string), Valid: true}
}

func ToSnakeCase(v string) string {
	snake := matchFirstCap.ReplaceAllString(v, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
