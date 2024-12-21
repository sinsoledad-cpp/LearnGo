package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	// "strconv"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:username`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data params should be a pointer")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data params should be a struct pointer")
		return
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	lineSlice := strings.Split(string(b), "\r\n")
	fmt.Printf("%#v\n", lineSlice)
	var structName string
	for idx, line := range lineSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax", idx+1)
				return
			}

			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])

			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i)
				fileType = filed
				if filed.Tag.Get("ini") == key {
					fieldName = filed.Name
					break
				}
			}
			if len(fileName) == 0 {
				continue
			}
			fileObj := sValue.FieldByName(fieldName)
			fmt.Println(fieldName, fileType.Type.Kind())
			switch fileType.Type.Kind() {
			// case reflect.String:
			// 	fileObj.SetString(value)

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
				// case reflect.Bool:
				// 	var valueBool bool
				// 	valueBool, err = strconv.ParseBool(value)
				// 	if err != nil {
				// 		err = fmt.Errorf("line:%d value type error", idx+1)
				// 		return
				// 	}
				// 	fileObj.SetBool(valueBool)
				// case reflect.Float32,reflect.Float64:
				// 	var valueFloat float64
				// 	valueFloat, err = strconv.ParseFloat(value,64)
				// 	if err != nil {
				// 		err = fmt.Errorf("line:%d value type error", idx+1)
				// 		return
				// 	}
				// 	fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	var cfg Config
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
