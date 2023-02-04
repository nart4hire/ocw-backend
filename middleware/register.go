package middleware

import (
	"reflect"
)

func (app *AppMiddlewares) Register() []Middleware {
	reflectValue := reflect.ValueOf(app)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()
	collections := []Middleware{}

	middlewareName := []string{}

	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)
		handler, ok := field.Interface().(Middleware)

		if !ok {
			continue
		}

		middlewareName = append(middlewareName, reflectType.Field(i).Name)
		collections = append(collections, handler)
	}

	if len(middlewareName) > 0 {
		app.Logger.Info("Registered Middlewares:")
		for _, middleware := range middlewareName {
			app.Logger.Info("- " + middleware)
		}
	} else {
		app.Logger.Info("No middleware registered")
	}

	return collections
}
