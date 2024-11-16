package util

import (
	"net/http"
	"reflect"
	"strings"
)

func ValidateFields[ReqModel any](obj ReqModel) ([]string, string) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	missingFields := []string{}

	// Iterate through all fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		jsonTag := fieldType.Tag.Get("json")

		// Skip fields without a JSON tag
		if jsonTag == "" {
			continue
		}

		// Split the JSON tag to handle cases where multiple tags are provided
		jsonTags := strings.Split(jsonTag, ",")
		jsonFieldName := jsonTags[0]

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				// Field is nil
				missingFields = append(missingFields, jsonFieldName)
			} else {
				// Field is not nil, check the value it points to
				fieldValue := field.Elem()
				if fieldValue.Kind() == reflect.String && fieldValue.String() == "" {
					// Field is a string and is empty
					missingFields = append(missingFields, jsonFieldName)
				}
			}
		} else if field.Kind() == reflect.String && field.String() == "" {
			// Handle non-pointer string fields if needed
			missingFields = append(missingFields, jsonFieldName)
		}
	}

	missingFieldsJsonErrMsg := "The following field(s) are missing or empty: " + strings.Join(missingFields, ", ")
	return missingFields, missingFieldsJsonErrMsg
}

func ExtractRequestBody[ReqBody any](r *http.Request) ReqBody {
	ctx := r.Context()
	val := ctx.Value("reqBodyObj")
	return val.(ReqBody)
}
