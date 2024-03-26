package util

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

const (
	B  int64 = 1 << (iota * 10) // iota = 0，左移0位，等于1^1 * 2^0 = 1 B
	KB                          // iota = 1，左移10位，等于1^1 * 2^10 = 1^11 = 1024 B = 1 KB
	MB
	GB
	TB
	PB
)

// ParseStrSize 解析表示数据大小的字符串，并将其转换为int64类型的字节数。
// 支持的单位包括KB、MB、GB、TB和PB，不区分大小写。
//
// 参数:
//
//	size - 需要解析的字符串，格式要求为“1KB”，“1MB”，“1GB”，“1TB”或“1PB”。
//
// 返回值:
//
//	int64 - 解析后的字节数表示，转换失败则返回0。
//	string - 原始传入的字符串大小。
//
// 注意：注释由openai ChatGPT4.0编写。
func ParseStrSize(size string) (int64, string, error) {
	// 将输入统一转换为大写，便于处理
	sizeUpper := strings.ToUpper(size)
	// 提取数字和单位
	var unit string
	var number int64
	_, err := fmt.Sscanf(sizeUpper, "%d%s", &number, &unit)
	if err != nil {
		log.Printf("传入格式错误,格式要求为“1KB”，“1MB”，“1GB”，“1TB”或“1PB”。")
		return 0, size, err
	}

	// 根据单位转换数字到对应的字节数
	switch unit {
	case "B":
		return number * B, size, nil
	case "KB":
		return number * KB, size, nil
	case "MB":
		return number * MB, size, nil
	case "GB":
		return number * GB, size, nil
	case "TB":
		return number * TB, size, nil
	case "PB":
		return number * PB, size, nil
	default:
		// 如果单位不识别，返回0和原始字符串
		log.Printf("未能识别的格式,采用默认参数100MB，格式要求为“1KB”，“1MB”，“1GB”，“1TB”或“1PB”。")
		return 100 * MB, size, nil
	}
}

func GetTypeSize(v interface{}) int64 {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	size := int64(val.Type().Size())

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			size += GetTypeSize(field.Interface()) // 递归处理结构体字段
		}
	} else if val.Kind() == reflect.Slice {
		elemSize := int64(val.Type().Elem().Size())
		size += int64(val.Len()) * elemSize
	} else if val.Kind() == reflect.String {
		size += int64(len(val.String()))
	}
	return size
}
