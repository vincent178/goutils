package di

import (
	"errors"
	"log"
	"reflect"
	"sync"
)

var container *Container

type Initializer interface {
	Init() error
}

type Closer interface {
	Close() error
}

type Container struct {
	services map[string]interface{}
}

func Load(data interface{}) error {
	rtype := reflect.TypeOf(data)
	rv := reflect.ValueOf(rtype)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("not a pointer")
	}

	name := rtype.Elem().Name()

	service := container.services[name]

	if service != nil {
		data = service
	} else {
		var once sync.Once
		once.Do(func() {
			if initializer, ok := data.(Initializer); ok {
				if err := initializer.Init(); err != nil {
					log.Fatal("init error", err)
				}
			}
			container.services[name] = data
		})
	}

	return nil
}

func Close() error {
	for _, service := range container.services {
		if closer, ok := service.(Closer); ok {
			if err := closer.Close(); err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	var once sync.Once
	once.Do(func() {
		container = &Container{
			services: map[string]interface{}{},
		}
	})
}

