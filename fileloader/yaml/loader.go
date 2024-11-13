package yaml

import (
	"errors"

	"gopkg.in/yaml.v3"
)

type Loader struct{}

const Name = "yaml"

func (Loader) Load(file []byte, config interface{}) (string, error) {
	err := yaml.Unmarshal(file, config)
	if err != nil {
		return Name, errors.New("load yaml file error")
	}
	return Name, nil
}
