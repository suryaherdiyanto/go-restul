package helper

import "database/sql"

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
