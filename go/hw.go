// Copyright 2016, yubo.
// All rights reserved.

// 一个例子,用来展示golang的基本特性
package hw // import "github.com/yubo/hw/go"

import (
	"fmt"
	"reflect"
)

const (
	h = "hello"
	w = "world"
)

// 输出 "hello,world"
func Hw() string {
	return fmt.Sprintf("%s,%s", h, w)
}

// 输出 "hello,world" + interface
func HwI(i interface{}) string {

	switch f := i.(type) {
	case bool:
		return fmt.Sprintf("%s %s %v", Hw(), "bool", f)
	case float32:
		return fmt.Sprintf("%s %s %v", Hw(), "float32", f)
	case float64:
		return fmt.Sprintf("%s %s %v", Hw(), "float64", f)
	case complex64:
		return fmt.Sprintf("%s %s %v", Hw(), "complex64", f)
	case complex128:
		return fmt.Sprintf("%s %s %v", Hw(), "complex128", f)
	case int:
		return fmt.Sprintf("%s %s %v", Hw(), "int", f)
	case int8:
		return fmt.Sprintf("%s %s %v", Hw(), "int8", f)
	case int16:
		return fmt.Sprintf("%s %s %v", Hw(), "int16", f)
	case int32:
		return fmt.Sprintf("%s %s %v", Hw(), "int32", f)
	case int64:
		return fmt.Sprintf("%s %s %v", Hw(), "int64", f)
	case uint:
		return fmt.Sprintf("%s %s %v", Hw(), "uint", f)
	case uint8:
		return fmt.Sprintf("%s %s %v", Hw(), "uint8", f)
	case uint16:
		return fmt.Sprintf("%s %s %v", Hw(), "uint16", f)
	case uint32:
		return fmt.Sprintf("%s %s %v", Hw(), "uint32", f)
	case uint64:
		return fmt.Sprintf("%s %s %v", Hw(), "uint64", f)
	case uintptr:
		return fmt.Sprintf("%s %s %v", Hw(), "uintptr", f)
	case string:
		return fmt.Sprintf("%s %s %v", Hw(), "string", f)
	case []byte:
		return fmt.Sprintf("%s %s %v", Hw(), "[]byte", f)
	case reflect.Value:
		return fmt.Sprintf("%s %s %v", Hw(), "reflect.Value", f)
	default:

	}
	return ""
}
