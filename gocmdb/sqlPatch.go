package main

import (
	"reflect"
	"strings"
)

const tagname = "sql"

type SQLPatch struct {
	Fields []string
	Args   []interface{}
}

func SQLPatches(resource interface{}) SQLPatch {
	var sqlPatch SQLPatch
	rType := reflect.TypeOf(resource)
	rVal := reflect.ValueOf(resource)
	n := rType.NumField()

	sqlPatch.Fields = make([]string, 0, n)
	sqlPatch.Args = make([]interface{}, 0, n)

	for i := 0; i < n; i++ {
		fType := rType.Field(i)
		fVal := rVal.Field(i)
		tag := fType.Tag.Get(tagname)

		// skip nil properties (not going to be patched), skip unexported fields, skip fields to be skipped for SQL
		if fVal.IsNil() || fType.PkgPath != "" || tag == "-" {
			continue
		}

		// if no tag is set, use the field name
		if tag == "" {
			tag = fType.Name
		}
		// and make the tag lowercase in the end
		tag = strings.ToLower(tag)

		sqlPatch.Fields = append(sqlPatch.Fields, tag+" = ?")

		var val reflect.Value
		if fVal.Kind() == reflect.Ptr {
			val = fVal.Elem()
		} else {
			val = fVal
		}

		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			sqlPatch.Args = append(sqlPatch.Args, val.Int())
		case reflect.String:
			sqlPatch.Args = append(sqlPatch.Args, val.String())
		case reflect.Bool:
			if val.Bool() {
				sqlPatch.Args = append(sqlPatch.Args, 1)
			} else {
				sqlPatch.Args = append(sqlPatch.Args, 0)
			}
		}
	}

	return sqlPatch
}
