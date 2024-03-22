package unmarshal

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/utils"
	"github.com/stevan-sdk/internal/validation"
)

type candidate struct {
	obj           any
	valid         bool
	optionalCount int
}

func ToComplexObject[T any](r *httptransport.Response) (*T, error) {
	result := new(T)
	err := unmarshalIntoProps(r.Body, result)
	if err != nil {
		return nil, err
	}

	candidates := createCandidatesFromProps(result)
	chosenCandidateIndex := chooseCandidateIndex(candidates)
	if chosenCandidateIndex == -1 {
		return nil, errors.New("cannot unmarshal response, no valid candidate found")
	}
	removeOtherCandidates(result, chosenCandidateIndex)

	return result, nil
}

// Try to Unmarshal the input data into the properties of a given struct.
func unmarshalIntoProps(data []byte, obj any) error {
	types := reflect.TypeOf(obj).Elem()
	values := reflect.ValueOf(obj).Elem()

	for i := 0; i < types.NumField(); i++ {
		fieldType := types.Field(i)
		unmarshalledValue := reflect.New(fieldType.Type)
		err := json.Unmarshal(data, unmarshalledValue.Interface())
		if err != nil {
			return err
		}

		value := unmarshalledValue.Elem()
		values.Field(i).Set(value)
	}

	return nil
}

func createCandidatesFromProps(obj any) []candidate {
	values := utils.GetReflectValue(reflect.ValueOf(obj))
	types := utils.GetReflectType(reflect.TypeOf(obj))

	candidates := make([]candidate, 0)
	for i := 0; i < types.NumField(); i++ {
		fieldValue := values.Field(i)

		value := fieldValue.Interface()
		candidate := candidate{
			obj:           value,
			valid:         isValid(value),
			optionalCount: countOptionals(value),
		}

		candidates = append(candidates, candidate)
	}

	return candidates
}

func countOptionals(candidate any) int {
	values := utils.GetReflectValue(reflect.ValueOf(candidate))
	types := utils.GetReflectType(reflect.TypeOf(candidate))

	count := 0
	for i := 0; i < types.NumField(); i++ {
		fieldValue := values.Field(i)
		fieldType := types.Field(i)

		if fieldValue.IsNil() {
			continue
		}

		if isOptional(fieldValue, fieldType) {
			count++
		}

		kind := utils.GetReflectKind(fieldType.Type)
		if kind == reflect.Struct || kind == reflect.Array || kind == reflect.Slice {
			count += countOptionals(fieldValue.Interface())
		}
	}

	return count
}

func isValid(candidate any) bool {
	err := validation.ValidateData(candidate)
	return err == nil
}

func isOptional(fieldValue reflect.Value, fieldType reflect.StructField) bool {
	required, found := fieldType.Tag.Lookup("required")
	return !found || required == "" || required == "false"
}

func chooseCandidateIndex(candidates []candidate) int {
	chosenCandidateIndex := -1
	chosenCandidateOptionalCount := -1

	for i, candidate := range candidates {
		if candidate.valid && candidate.optionalCount > chosenCandidateOptionalCount {
			chosenCandidateIndex = i
			chosenCandidateOptionalCount = candidate.optionalCount
		}
	}

	return chosenCandidateIndex
}

func removeOtherCandidates(obj any, chosenCandidateIndex int) {
	values := utils.GetReflectValue(reflect.ValueOf(obj))
	types := utils.GetReflectType(reflect.TypeOf(obj))

	for i := 0; i < types.NumField(); i++ {
		if i != chosenCandidateIndex {
			fieldValue := values.Field(i)
			fieldValue.Set(reflect.Zero(fieldValue.Type()))
		}
	}
}
