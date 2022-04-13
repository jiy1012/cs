package yaml

import (
	"errors"
	"gopkg.in/yaml.v2"
)

type Loader struct{}

func (Loader) Load(file []byte, config interface{}) error {
	err := yaml.Unmarshal(file, config)
	if err != nil {
		return errors.New("load yaml file error")
	}
	return nil
}
