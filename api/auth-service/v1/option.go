package v1

import (
	"net/url"
	"strings"
)

// SubjectObjectsListOptions 定义获取获取访问者拥有的资源 query 参数
type SubjectObjectsListOptions struct {
	// 访问者类型
	SubjectType SubjectType `json:"subject_type,omitempty"`
	// 操作者 ID
	SubjectID string `json:"subject_id,omitempty"`
	// 资源类型列表
	ObjectTypes []ObjectType `json:"object_type,omitempty"`
}

// MarshalQueryParameter 将 SubjectObjectsListOptions 转为 query 字符串
func (opts *SubjectObjectsListOptions) MarshalQueryParameter() (string, error) {
	v := make(url.Values)

	// subject_type
	if opts.SubjectType != "" {
		v.Set("subject_type", string(opts.SubjectType))
	}
	// subject_id
	if opts.SubjectID != "" {
		v.Set("subject_id", opts.SubjectID)
	}
	// object_type
	var objectTypes []string
	for _, t := range opts.ObjectTypes {
		objectTypes = append(objectTypes, string(t))
	}
	if objectTypes != nil {
		v.Set("object_type", strings.Join(objectTypes, ","))
	}

	return v.Encode(), nil
}

// IndicatorAuthorizingRequestGetOptions 定义获取指标授权申请的 query 参数
type IndicatorAuthorizingRequestGetOptions struct {
	// 是否获取指标授权申请所引用的资源
	Reference bool `json:"reference,omitempty"`
}

// APIAuthorizingRequestGetOptions 定义获取接口授权申请的 query 参数
type APIAuthorizingRequestGetOptions struct {
	// 是否获取接口授权申请所引用的资源
	Reference bool `json:"reference,omitempty"`
}
