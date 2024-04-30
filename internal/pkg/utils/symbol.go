package utils

import "github.com/golang/protobuf/protoc-gen-go/generator"

// 返回骆驼峰
func CamelCase(s string) string {
	return generator.CamelCase(s)
}
