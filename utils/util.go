package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

// Ucfirst 首字母大写
func Ucfirst(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToUpper(str[0:1]) + str[1:]
}

// Lcfirst 首字母小写
func Lcfirst(str string) string {
	if len(str) == 0 {
		return str
	}
	return strings.ToLower(str[0:1]) + str[1:]
}

func Case2CamelSpecial(name string, upperFirst bool, chars []string) string {
	for _, s := range chars {
		name = strings.Replace(name, s, " ", -1)
	}
	arr := strings.Split(name, " ")
	for i, s := range arr {
		arr[i] = strings.ToUpper(s[0:1]) + s[1:]
	}
	name = strings.Join(arr, "")
	if upperFirst {
		return name
	}
	return strings.ToLower(name[0:1]) + name[1:]
}

func LoadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return b, nil
}

//写入文件
func WriteFile(filename string, content []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return nil
}
