package helpers

import (
	"encoding/json"

	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
)

type JsonHelper struct {
}

// Marshal implements repository.JsonRepository
func (*JsonHelper) Marshal(i interface{}, v interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

// Unmarshal implements repository.JsonRepository
func (*JsonHelper) Unmarshal(item interface{}) error {
	panic("unimplemented")
}

func NewJsonHelper() repository.JsonRepository {
	return &JsonHelper{}
}
