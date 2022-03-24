package helpers

import (
	"encoding/json"

	"github.com/negiseijin/clean-architectur-swagger/domain/repository"
)

type Helpers struct {
}

// JsonToStruct implements repository.HelpersRepository
func (*Helpers) JsonToStruct(i interface{}, v interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

func NewHelpers() repository.HelpersRepository {
	return &Helpers{}
}
