package maps

import (
	"reflect"
	"strconv"

	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

func BuildFormString(struct_data interface{}) (map[string]string, error) {
	form := map[string]string{}

	valueOfStruct := reflect.ValueOf(struct_data)

	// Verificar si el argumento es una estructura
	if valueOfStruct.Kind() != reflect.Struct {
		return nil, model.Error("el argumento debe ser una estructura")
	}
	// Obtener el tipo de la estructura
	structType := valueOfStruct.Type()
	// struct_name := structType.Name()

	structValue := reflect.ValueOf(struct_data)

	for i := 0; i < structValue.NumField(); i++ {
		name_value := structType.Field(i).Name

		// solo si la primera letra es mayúscula
		if _, ok := strings.VALID_LETTERS[name_value[0]]; ok {

			field := structValue.Field(i)
			fieldType := field.Type()

			name_value = strings.LowerCaseFirstLetter(name_value)

			// Convertir valores a string según el tipo
			var fieldValue string
			switch fieldType.Kind() {
			case reflect.String:
				fieldValue = field.Interface().(string)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fieldValue = strconv.FormatInt(field.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fieldValue = strconv.FormatUint(field.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				fieldValue = strconv.FormatFloat(field.Float(), 'f', -1, 64)
			case reflect.Bool:
				fieldValue = strconv.FormatBool(field.Bool())
			default:
				// continuar si el otro valor no soportado
				continue
			}

			form[name_value] = fieldValue

		}
	}

	return form, nil
}
