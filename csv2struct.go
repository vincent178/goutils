package goutils

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	errNotImplement = errors.New("not implement")
)

type option struct {
	SuppressError   bool
	CaseInsensitive bool
}

func WithSuppressError(suppressError bool) func(*option) {
	return func(o *option) {
		o.SuppressError = suppressError
	}
}

func WithCaseInsensitive(caseInsensitive bool) func(*option) {
	return func(o *option) {
		o.CaseInsensitive = caseInsensitive
	}
}

// CsvMapToStruct use csv tag unmarshal to struct
func CsvMapToStruct(src map[string]string, out interface{}, options ...func(*option)) error {
	var o option
	for _, f := range options {
		f(&o)
	}

	rv := reflect.ValueOf(out)

	if rv.IsNil() {
		return errors.New("invalid nil out")
	}

	if !(rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) {
		return errors.New("invalid type")
	}

	e := rv.Elem()

	if reflect.TypeOf(e).Kind() != reflect.Struct {
		return errors.New("invalid type")
	}

	for i := 0; i < e.NumField(); i++ {
		// get map key and value
		key := e.Type().Field(i).Tag.Get("csv")

		var val string
		for k, v := range src {
			if o.CaseInsensitive {
				if strings.EqualFold(k, key) {
					val = v
					break
				}
			} else {
				if k == key {
					val = v
					break
				}
			}
		}

		if !e.Field(i).CanSet() {
			continue
		}

		// handle empty value
		if val == "" {
			e.Field(i).Set(e.Field(i))
			continue
		}

		switch e.Type().Field(i).Type.Kind() {
		case reflect.String:
			e.Field(i).SetString(val)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				if o.SuppressError {
					// set default value
					e.Field(i).Set(e.Field(i))
					continue
				}
				return err
			}
			e.Field(i).SetUint(x)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				if o.SuppressError {
					// set default value
					e.Field(i).Set(e.Field(i))
					continue
				}
				return err
			}
			e.Field(i).SetInt(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(val, 64)
			if err != nil {
				if o.SuppressError {
					// set default value
					e.Field(i).Set(e.Field(i))
					continue
				}
				return err
			}
			e.Field(i).SetFloat(x)
		case reflect.Bool:
			x, err := strconv.ParseBool(val)
			if err != nil {
				if o.SuppressError {
					// set default value
					e.Field(i).Set(e.Field(i))
					continue
				}
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
