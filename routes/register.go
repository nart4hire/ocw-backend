package routes

import (
	"reflect"
)

func (app *AppRouter) Register() []RouteGroup {
	reflectValue := reflect.ValueOf(app)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	collections := []RouteGroup{}

	for i := 0; i < reflectValue.NumField(); i++ {
		handler, ok := reflectValue.Field(i).Interface().(RouteGroup)

		if !ok {
			continue
		}

		collections = append(collections, handler)
  }

	return collections
}
