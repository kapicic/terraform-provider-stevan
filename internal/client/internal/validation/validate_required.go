package validation

import (
	"fmt"
	"reflect"
)

func validateRequired(fieldValue reflect.Value, fieldType reflect.StructField) error {
	required, found := fieldType.Tag.Lookup("required")
	if !found || required == "" || required == "false" {
		return nil
	}

	if fieldValue.IsNil() {
		return fmt.Errorf("field %s is required", fieldType.Name)
	}

	return nil
}
