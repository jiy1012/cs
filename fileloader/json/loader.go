package json

import (
	"encoding/json"
	"errors"
)

type Loader struct{}

const Name = "json"

func (Loader) Load(file []byte, config interface{}) (string, error) {
	err := json.Unmarshal(file, config)
	if err != nil {
		return Name, errors.New("load json file error")
	}
	return Name, nil
}
