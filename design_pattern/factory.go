package design_pattern

import (
	"fmt"
	"strings"
)

type AbsParser interface {
	Parse(string) string
}

type JsonParser struct{}

func (_ *JsonParser) Parse(msg string) (res string) {
	res = strings.ToUpper(msg)
	fmt.Println(res)
	return
}

type YamlParser struct{}

func (_ *YamlParser) Parse(msg string) (res string) {
	msgByte := []byte(msg)
	msgByte = append(msgByte, "hello"...)
	res = string(msgByte)
	fmt.Println(res)
	return
}

type ParseFactory struct{}

func (_ *ParseFactory) NewFactory(t string) AbsParser {
	switch t {
	case "json":
		return &JsonParser{}
	case "yaml":
		return &YamlParser{}
	default:
		return nil
	}
}
