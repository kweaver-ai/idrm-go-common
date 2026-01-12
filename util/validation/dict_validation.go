package validation

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/kweaver-ai/idrm-go-common/rest/base"
	driven "github.com/kweaver-ai/idrm-go-common/rest/configuration_center"
)

// StructField 定义字段信息结构体
type StructField struct {
	Name      string            // 字段名
	Type      string            // 字段类型
	Tags      map[string]string // 标签键值对
	IsPointer bool              // 是否为指针类型
	Value     interface{}       // 字段值
	Children  []StructField     // 嵌套结构体字段
}

// GetStructInfo 获取结构体的所有字段信息
func GetStructInfo(s interface{}) ([]StructField, error) {
	if s == nil {
		return nil, fmt.Errorf("input struct cannot be nil")
	}

	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // 获取指针指向的值
	}

	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct, got %s", v.Kind())
	}

	return getFieldsInfo(v), nil
}

// getFieldsInfo 递归获取字段信息
func getFieldsInfo(v reflect.Value) []StructField {
	t := v.Type()
	fields := make([]StructField, 0, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// 过滤私有字段和空值字段（未导出字段）
		if !field.IsExported() || value.Interface() == nil {
			continue
		}

		structField := StructField{
			Name:      field.Name,
			Type:      field.Type.String(),
			Tags:      parseTagsToMap(string(field.Tag)),
			IsPointer: field.Type.Kind() == reflect.Ptr,
		}

		// 获取字段值
		if structField.IsPointer && !value.IsNil() {
			structField.Value = value.Interface() // 获取指针指向的值
		} else {
			structField.Value = value.Interface() // 获取字段值
		}

		// 处理嵌套结构体和切片
		switch field.Type.Kind() {
		case reflect.Struct:
			if childFields := getFieldsInfo(value); childFields != nil {
				structField.Children = childFields
			}
		case reflect.Slice:
			structField.Children = handleSlice(value)
		case reflect.Ptr:
			if field.Type.Elem().Kind() == reflect.Struct && !value.IsNil() {
				if childFields := getFieldsInfo(value.Elem()); childFields != nil {
					structField.Children = childFields
				}
			}
		}

		fields = append(fields, structField)
	}

	return fields
}

// handleSlice 处理切片类型
func handleSlice(value reflect.Value) []StructField {
	var children []StructField
	elemKind := value.Type().Elem().Kind()

	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		if elemKind == reflect.Ptr && !elem.IsNil() {
			if childFields := getFieldsInfo(elem.Elem()); childFields != nil {
				children = append(children, childFields...)
			}
		} else if elemKind == reflect.Struct {
			if childFields := getFieldsInfo(elem); childFields != nil {
				children = append(children, childFields...)
			}
		}
	}
	return children
}

// parseTagsToMap 解析标签字符串为map
func parseTagsToMap(tag string) map[string]string {
	tags := make(map[string]string)
	tag = strings.Trim(tag, "`")
	if tag == "" {
		return tags
	}

	for _, tagPair := range strings.Split(tag, " ") {
		if tagPair == "" {
			continue
		}

		parts := strings.SplitN(tagPair, ":", 2)
		if len(parts) == 2 {
			key := strings.Trim(parts[0], "\"")
			value := strings.Trim(parts[1], "\"")
			tags[key] = value
		}
	}

	return tags
}

func CheckDictTypeKey(ctx context.Context, s interface{}) (string, error) {
	// 获取结构体信息
	fields, err := GetStructInfo(s)
	if err != nil {
		return "", err
	}
	req, _ := printStructFields(fields)
	if req.DictTypeKey == nil || len(req.DictTypeKey) == 0 {
		return "", err
	}
	//获取接口
	resp, err := batchCheckNotExistTypeKey(ctx, req)
	if err != nil {
		return "", err
	}
	str := strings.Join(resp, "，")
	return str, err
}

// 获取form或者json字段、数据类型字段、数据值
func printStructFields(fields []StructField) (driven.CheckDictTypeKeyReq, error) {
	// printStructFields 处理字段信息并返回一个新的 Req 对象
	req := driven.CheckDictTypeKeyReq{} // 创建一个新的 Req 对象
	for _, field := range fields {
		// 处理当前字段
		tagDictType := field.Tags["DictType"]
		if tagDictType != "" {
			tagForm := field.Tags["form"]
			if tagForm == "" {
				tagForm = field.Tags["json"]
			}
			req.DictTypeKey = append(req.DictTypeKey, &driven.DictTypeKey{
				DictType: tagDictType,
				Field:    tagForm,
				DictKey:  fmt.Sprintf("%v", field.Value),
			})
		}

		// 处理子字段
		if len(field.Children) > 0 {
			// 递归处理子字段并将结果合并到 req 中
			childReq, err := printStructFields(field.Children)
			if err != nil {
				return req, err
			}
			req.DictTypeKey = append(req.DictTypeKey, childReq.DictTypeKey...) // 合并子请求
		}
	}

	return req, nil
}

// defaultHTTPClient 返回一个配置好的HTTP客户端
func defaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
			IdleConnTimeout:     30 * time.Second,
			DisableCompression:  false,
		},
		Timeout: 3 * time.Second,
	}
}
func batchCheckNotExistTypeKey(ctx context.Context, req driven.CheckDictTypeKeyReq) ([]string, error) {
	client := defaultHTTPClient()
	url := base.Service.ConfigurationCenterHost + "/api/internal/configuration-center/v1/dict/batch-check-type-key"

	items, err := base.Call[[]string](ctx, client, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}
	resp := FindDifference(req.DictTypeKey, items)
	return resp, err
}

// findDifference 返回在 array1 中但不在 array2 中的字符串的字段名称
func FindDifference(array1 []*driven.DictTypeKey, array2 []string) []string {
	m := make(map[string]struct{})
	n := make(map[string]struct{})
	// 将 array2 中的元素放入 m 中
	for _, v := range array2 {
		m[v] = struct{}{}
	}

	// 遍历 array1，找出在 m 中的元素，即 array1、array2 共有的元素
	for _, v := range array1 {
		if _, ok := m[v.DictType]; ok {
			n[v.Field] = struct{}{}
		}
	}
	// 将 map n 的 key 转换为 slice
	diff := make([]string, len(n))
	i := 0
	for k := range n {
		diff[i] = k
		i++
	}
	return diff
}
