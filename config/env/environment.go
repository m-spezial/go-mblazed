package env

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ReadValues loads the annotation from the given struct and replace the values with the environment variables.
// The struct must be a pointer to a struct.
//
func ReadValues(value interface{}) error {
	ps := reflect.ValueOf(value)
	t := ps.Elem()

	if t.Kind() != reflect.Struct {
		return errors.New("object need to be a struct")
	}
	elem := reflect.TypeOf(value).Elem()

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := elem.Field(i).Tag.Get("env")

		if len(tag) == 0 || tag == "-" {
			continue
		}

		parts := strings.Split(tag, ",")
		required := false

		if len(parts) > 1 && parts[1] == "required" {
			required = true
		}

		value := os.Getenv(parts[0])

		if value == "" {

			if required {
				return errors.New("variable for env variable '" + parts[0] + "' is required")

			} else if len(parts) > 1 && strings.HasPrefix(parts[1], "default") {
				defaultSplit := strings.SplitN(parts[1], "=", 2)

				if len(defaultSplit) == 1 {
					return errors.New("default not give a value with syntax default=value (" + parts[1] + ")")
				}

				value = defaultSplit[1]
			} else {
				continue
			}
		}

		if field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Int64:
				fallthrough
			case reflect.Int32:
				fallthrough
			case reflect.Int16:
				fallthrough
			case reflect.Int8:
				fallthrough
			case reflect.Int:
				intVal, err := strconv.ParseInt(value, 10, 64)

				if err != nil {
					return err
				}

				field.SetInt(intVal)
			case reflect.Bool:
				boolVal, err := strconv.ParseBool(value)

				if err != nil {
					return err
				}

				field.SetBool(boolVal)
			case reflect.Uint64:
				fallthrough
			case reflect.Uint32:
				fallthrough
			case reflect.Uint16:
				fallthrough
			case reflect.Uint8:
				fallthrough
			case reflect.Uint:
				intVal, err := strconv.ParseUint(value, 10, 64)

				if err != nil {
					return err
				}

				field.SetUint(intVal)
			}

		}
	}
	return nil
}
