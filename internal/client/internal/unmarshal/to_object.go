package unmarshal

import (
	"encoding/json"

	"github.com/stevan-sdk/internal/clients/rest/httptransport"
)

func ToObject[T any](r *httptransport.Response) (*T, error) {
	result := new(T)
	err := json.Unmarshal(r.Body, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
