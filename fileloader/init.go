package fileloader

import (
	"github.com/jiy1012/cs/fileloader/json"
	"github.com/jiy1012/cs/fileloader/toml"
	"github.com/jiy1012/cs/fileloader/yaml"
)

var SupportedExts = []string{"json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"}
var LoaderRegistrys = NewLoaderRegistry()

func init() {
	{
		l := json.Loader{}
		LoaderRegistrys.RegisterDecoder("json", l)
	}
	{
		l := yaml.Loader{}
		LoaderRegistrys.RegisterDecoder("yaml", l)
		LoaderRegistrys.RegisterDecoder("yml", l)
	}
	{
		l := toml.Loader{}
		LoaderRegistrys.RegisterDecoder("toml", l)
	}
}
