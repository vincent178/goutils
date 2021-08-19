package goutils

import (
	"errors"
	"reflect"
	"strconv"
)

/*
CsvMapToStruct use csv tag

type Person struct {
	Name string `csv:"Name"`
	Age int     `csv:"Age"`
}
*/
func CsvMapToStruct(src map[string]string, out interface{}) error {
	rv := reflect.ValueOf(out)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("invalid type")
	}

	e := rv.Elem()

	if reflect.TypeOf(e).Kind() != reflect.Struct {
		return errors.New("invalid type")
	}

	for i := 0; i < e.NumField(); i++ {
		// get map key and value
		key := e.Type().Field(i).Tag.Get("csv")
		val := src[key]

		switch e.Type().Field(i).Type.Kind() {
		case reflect.String:
			e.Field(i).SetString(val)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			num, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			e.Field(i).SetInt(int64(num))
		}
	}

	return nil
}
