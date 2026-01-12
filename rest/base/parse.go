package base

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"unicode"

	"github.com/kweaver-ai/idrm-go-common/interception"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
)

const EnginePathPattern = "/api/virtual_engine_service"

const (
	ParamTypeStructTag = "param_type"

	ParamKeyTypeUri   = "uri"
	ParamKeyTypeQuery = "query"
	ParamKeyTypeForm  = "form"
	ParamKeyTypeJson  = "json"
)

type EmptyArgs struct{}

// NewRequest  根据参数，返回指定的
func NewRequest(ctx context.Context, method, path string, args any) (*http.Request, error) {
	value := reflect.ValueOf(args)

	for value.Kind() == reflect.Pointer {
		if value.IsNil() {
			value = reflect.New(value.Elem().Type())
		}
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct && value.Kind() != reflect.Slice {
		panic("req param T must struct slice or map")
	}

	queryValues := make(url.Values)
	formUrlEncoded := make(url.Values)
	bodyValues := make(map[string]any)
	bodyValueSlice := make([]any, 0)

	typ := value.Type()
	//需要通过反射设置的值
	bodyRef := reflect.ValueOf(bodyValues)

	// 默认这两种是json类型的
	if value.Kind() == reflect.Slice {
		for i := 0; i < value.Len(); i++ {
			bodyValueSlice = append(bodyValueSlice, value.Index(i).Interface())
		}
	} else {
		//只处理最外层参数
		for i := 0; i < typ.NumField(); i++ {
			fieldType := typ.Field(i)
			fieldValue := value.Field(i)

			// 忽略未导出的字段（以小写字母开头的字段）
			if !IsExported(fieldType) {
				continue
			}
			//处理路径参数
			if k, vs := getKeyStringValue(fieldType, fieldValue, ParamKeyTypeUri); len(vs) > 0 {
				path = strings.ReplaceAll(path, ":"+k, vs[0])
				continue
			}
			//处理query参数
			if k, vs := getKeyStringValue(fieldType, fieldValue, ParamKeyTypeQuery); len(vs) > 0 {
				queryValues[k] = vs
				continue
			}
			//处理form参数
			if k, vs := getKeyStringValue(fieldType, fieldValue, ParamKeyTypeForm); len(vs) > 0 {
				formUrlEncoded[k] = vs
				continue
			}
			//处理body参数
			if k, vs := getKeyJsonValue(fieldType, fieldValue, ParamKeyTypeJson); k != "" {
				bodyRef.SetMapIndex(reflect.ValueOf(k), vs)
				continue
			}
		}
	}

	var reader io.Reader
	header := make(http.Header)
	//处理query参数
	if len(queryValues) > 0 {
		path = path + "?" + queryValues.Encode()
	}
	//处理body参数
	if len(bodyValues) > 0 || len(bodyValueSlice) > 0 {
		if len(bodyValues) > 0 {
			bts, err := json.Marshal(bodyValues)
			if err != nil {
				return nil, err
			}
			fmt.Printf("base.NewRequest body:%v\n", string(bts))
			reader = bytes.NewReader(bts)
		}
		if len(bodyValueSlice) > 0 {
			bts, err := json.Marshal(bodyValueSlice)
			if err != nil {
				return nil, err
			}
			fmt.Printf("base.NewRequest body:%v\n", string(bts))
			reader = bytes.NewReader(bts)
		}
		header.Set("Content-Type", "application/json")
	}
	//处理form-data参数
	if len(formUrlEncoded) > 0 {
		reader = strings.NewReader(formUrlEncoded.Encode())
		header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	newReq, err := http.NewRequestWithContext(ctx, method, path, reader)
	if err != nil {
		return nil, err
	}
	//Content-Type
	contentType := header.Get("Content-Type")
	if contentType != "" {
		newReq.Header.Set("Content-Type", contentType)
	}
	if strings.Contains(path, EnginePathPattern) {
		newReq.Header.Set("X-Presto-User", "admin")
	}
	// 添加认证信息
	interception.SeAuthorizationIfEmpty(ctx, newReq.Header)
	return newReq, nil
}

func getFieldStringValue(fieldValue reflect.Value) []string {
	kind := fieldValue.Kind()
	switch {
	case kind == reflect.String:
		return []string{fieldValue.String()}

	case kind == reflect.Bool:
		return []string{fmt.Sprintf("%v", fieldValue.Bool())}

	case kind >= reflect.Int && kind <= reflect.Int64:
		return []string{fmt.Sprintf("%v", fieldValue.Int())}

	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return []string{fmt.Sprintf("%v", fieldValue.Uint())}

	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return []string{fmt.Sprintf("%v", fieldValue.Float())}

	case kind >= reflect.Complex64 && kind <= reflect.Complex128:
		return []string{fmt.Sprintf("%v", fieldValue.Complex())}

	case kind == reflect.Slice || kind == reflect.Array:
		if fieldValue.Len() <= 0 {
			return []string{}
		}
		ss := make([]string, 0)
		for i := 0; i < fieldValue.Len(); i++ {
			fieldElement := fieldValue.Index(i)
			ss = append(ss, getFieldStringValue(fieldElement)...)
		}
		return ss
	default:
		log.Warnf("unsupported type %s for struct field %s", fieldValue.Kind(), fieldValue.Type().Name())
	}
	return []string{}
}

func getKeyStringValue(fieldType reflect.StructField, fieldValue reflect.Value, tag string) (string, []string) {
	key := fieldType.Tag.Get(tag)
	if key == "" {
		return "", []string{}
	}
	return key, getFieldStringValue(fieldValue)
}

func getKeyJsonValue(fieldType reflect.StructField, fieldValue reflect.Value, tag string) (string, reflect.Value) {
	t := fieldType.Tag.Get(tag)
	if t == "-" {
		return "", reflect.Value{}
	}
	key, opts := parseTag(t)
	if opts.Contains("omitempty") && isEmptyValue(fieldValue) {
		return "", reflect.Value{}
	}
	return key, fieldValue
}

// IsExported 判断一个反射结构体字段是否可导出
func IsExported(field reflect.StructField) bool {
	// 如果字段属于当前包（PkgPath为空），并且首字母大写，则字段是可导出的
	return field.PkgPath == "" && unicode.IsUpper(rune(field.Name[0]))
}

// tagOptions is the string following a comma in a struct field's "json"
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

// parseTag splits a struct field's json tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// Contains reports whether a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Interface, reflect.Pointer:
		return v.IsZero()
	}
	return false
}
