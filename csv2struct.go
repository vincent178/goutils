package goutils

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	errNotImplement = errors.New("not implement")
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
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			e.Field(i).SetUint(x)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			e.Field(i).SetInt(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			e.Field(i).SetFloat(x)
		case reflect.Bool:
			x, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			e.Field(i).SetBool(x)
		case reflect.Complex64, reflect.Complex128:
		case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
			// TODO: handle Ptr
		case reflect.Ptr:
			return errNotImplement
		}
	}

	return nil
}
