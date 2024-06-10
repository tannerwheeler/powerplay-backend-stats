package validators

import (
	"fmt"
	"github.com/jak103/powerplay/internal/models"
	"reflect"
	"strings"
)

func IsValidSortField(field string, model interface{}) bool {
	modelValue := reflect.ValueOf(model)
	modelType := modelValue.Type()
	lowerField := strings.ToLower(field) // Convert field to lowercase
	fmt.Println("lowerField: ", lowerField)

	// Traverse through the fields of the struct
	for i := 0; i < modelType.NumField(); i++ {
		structField := modelType.Field(i)

		if structField.Anonymous && structField.Type == reflect.TypeOf(models.DbModel{}) {
			// Check fields in the embedded DbModel
			embeddedFieldValue := modelValue.Field(i)
			embeddedFieldType := embeddedFieldValue.Type()

			for j := 0; j < embeddedFieldType.NumField(); j++ {
				embeddedField := embeddedFieldType.Field(j)
				jsonTag := embeddedField.Tag.Get("json")
				jsonField := strings.Split(jsonTag, ",")[0]
				fmt.Println("embeddedField JSON: ", strings.ToLower(jsonField))
				if strings.ToLower(jsonField) == lowerField {
					return true
				}
			}
		} else {
			jsonTag := structField.Tag.Get("json")
			jsonField := strings.Split(jsonTag, ",")[0]
			fmt.Println("structField JSON: ", strings.ToLower(jsonField))
			if strings.ToLower(jsonField) == lowerField {
				return true
			}
		}
	}
	return false
}
