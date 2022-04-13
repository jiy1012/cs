package json

import (
	"encoding/json"
	"errors"
)

type Loader struct{}

func (Loader) Load(file []byte, config interface{}) error {
	err := json.Unmarshal(file, config)
	if err != nil {
		return errors.New("load json file error")
	}
	return nil
}
