package toml

import (
	"errors"

	"github.com/pelletier/go-toml"
)

type Loader struct{}

const Name = "toml"

func (Loader) Load(file []byte, config interface{}) (string, error) {
	tree, err := toml.LoadBytes(file)
	if err != nil {
		return Name, errors.New("load toml file error")
	}
	if m, ok := config.(*map[string]interface{}); ok {
		vmap := *m
		tmap := tree.ToMap()
		for k, v := range tmap {
			vmap[k] = v
		}
		return Name, nil
	}
	err = tree.Unmarshal(config)
	if err != nil {
		return Name, errors.New("load toml file error")
	}
	return Name, nil
}
