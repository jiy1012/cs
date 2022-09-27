package main

import (
	"errors"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/jiy1012/cs/fileloader"
	"github.com/jiy1012/cs/utils"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

var args Args

func main() {
	arg.MustParse(&args)
	//fmt.Println(args)
	cFile := args.Input
	ext := strings.TrimLeft(filepath.Ext(cFile), ".")
	file := strings.TrimSuffix(filepath.Base(cFile), "."+ext)
	//fmt.Println(file)
	var c interface{}
	fileBytes, err := utils.LoadFile(cFile)
	if err != nil {
		fmt.Println("load file error:", err)
		os.Exit(1)
	}
	err = fileloader.LoaderRegistrys.Load(ext, fileBytes, &c)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	structName := utils.Case2CamelSpecial(file, true, []string{"_", "-"})
	configType := reflect.TypeOf(c)
	switch configType.Kind() {
	case reflect.Map:
		parseMap(c, structName)
	case reflect.Array:
	case reflect.Slice:
		if len(c.([]interface{})) > 0 {
			_ = parseSlice(c.([]interface{}), structName)
		}
	}
}

func parseSlice(c []interface{}, key string) string {
	vTppe := reflect.TypeOf(c[0]).Kind()
	v := reflect.ValueOf(c[0])
	fType := ""
	switch vTppe {
	case reflect.Int, reflect.Int8:
		fType = vTppe.String()
	case reflect.Int32, reflect.Int64:
		fType = vTppe.String()
	case reflect.String:
		fType = vTppe.String()
	case reflect.Map:
		parseMap(c[0], key)
		fType = key
	case reflect.Slice:
		if len(v.Interface().([]interface{})) > 0 {
			sType := parseSlice(v.Interface().([]interface{}), v.String())
			fType = "[]" + sType
		} else {
			fType = "[]interface{}"
		}
	case reflect.Bool:
		fType = vTppe.String()
	case reflect.Float64:
		fType = vTppe.String()
	}
	return fType
}

func getKeyString(key reflect.Value) (string, error) {
	switch key.Kind() {
	case reflect.String:
		return key.String(), nil
	case reflect.Interface:
		s, err := getInterfaceKey(key)
		return s, err
	}
	return "", errors.New(fmt.Sprintf("golang do not support type: %+v  for key: %+v ", key.Kind().String(), key))
}

func getInterfaceKey(key reflect.Value) (string, error) {
	t := reflect.TypeOf(key.Interface()).Kind()
	switch t {
	case reflect.String:
		return key.Interface().(string), nil
	}
	return "", errors.New(fmt.Sprintf("golang do not support type: %+v  for key: %+v ", t.String(), key))
}

func parseMap(c interface{}, structName string) {
	var m []KvStruct
	fKey, fType := "", ""
	iter := reflect.ValueOf(c).MapRange()
	for iter.Next() {
		k := iter.Key()
		kString, err := getKeyString(k)
		if err != nil {
			fmt.Println("get key error:", err)
			os.Exit(1)
		}
		//kType := reflect.TypeOf(k).Kind()
		//fmt.Println("ktype:", kString)
		v := iter.Value()
		if v.Interface() != nil {
			vTppe := reflect.TypeOf(v.Interface()).Kind()
			switch vTppe {
			case reflect.Int, reflect.Int8:
				fKey = utils.Ucfirst(kString)
				fType = vTppe.String()
			case reflect.Int32, reflect.Int64:
				fKey = utils.Ucfirst(kString)
				fType = vTppe.String()
			case reflect.String:
				fKey = utils.Ucfirst(kString)
				fType = vTppe.String()
			case reflect.Map:
				if args.AutoAddPerfix {
					fKey, fType = structName+utils.Ucfirst(kString), structName+utils.Ucfirst(kString)
				} else {
					fKey, fType = utils.Ucfirst(kString), utils.Ucfirst(kString)
				}
				parseMap(v.Interface(), fKey)
			case reflect.Slice:
				if len(v.Interface().([]interface{})) > 0 {
					sType := parseSlice(v.Interface().([]interface{}), utils.Ucfirst(kString))
					fKey, fType = utils.Ucfirst(kString), "[]"+sType
				} else {
					fKey = utils.Ucfirst(k.String())
					fType = "[]interface{}"
				}
			case reflect.Bool:
				fKey = utils.Ucfirst(kString)
				fType = vTppe.String()
			case reflect.Float64:
				fKey = utils.Ucfirst(kString)
				fType = vTppe.String()
			}
		} else {
			fKey = utils.Ucfirst(kString)
			fType = "interface{}"
		}

		m = append(m, KvStruct{
			Field:     fKey,
			FieldType: fType,
		})
	}
	WriteStruct(structName, m)
	return
}

func WriteStruct(structName string, m []KvStruct) {
	fileContent := "package " + args.Package + "\n"
	fileContent += "type " + structName + " struct { \n"
	for _, kvStruct := range m {
		fileContent += kvStruct.Field + " " + kvStruct.FieldType + "\n"
	}
	fileContent += "}\n"
	if args.Output != "" {
		if _, err := os.Stat(args.Output); os.IsNotExist(err) {
			os.MkdirAll(args.Output, os.ModePerm)
		}
	}
	fileName := filepath.Join(args.Output, structName+".go")
	if _, err := os.Stat(fileName); err == nil || !os.IsNotExist(err) {
		dErr := os.Remove(fileName)
		if dErr != nil {
			// 删除失败
			fmt.Println("delete old file error :", dErr)
		}
	}
	e := utils.WriteFile(fileName, []byte(fileContent))
	if args.GoRoot != "" {
		goFmt := filepath.Join(args.GoRoot, "bin", "gofmt")
		err := exec.Command(goFmt, "-l", "-w", fileName).Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if e != nil {
		fmt.Println("write file error:", e)
		os.Exit(1)
	}
}
