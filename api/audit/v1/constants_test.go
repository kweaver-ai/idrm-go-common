package v1

import (
	"reflect"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestOperation_SimplifiedChineseName(t *testing.T) {
	tests := []struct {
		name string
		o    Operation
		want string
	}{
		{
			name: "未定义的操作",
			o:    "Hello Kitty",
			want: "未知操作",
		},
		{
			name: "生成接口",
			o:    OperationGenerateAPI,
			want: "生成接口",
		},
		{
			name: "注册接口",
			o:    OperationRegisterAPI,
			want: "注册接口",
		},
		{
			name: "修改接口",
			o:    OperationUpdateAPI,
			want: "修改接口",
		},
		{
			name: "设置接口权限",
			o:    OperationSetAPIAuthorization,
			want: "设置接口权限",
		},
		{
			name: "发布接口",
			o:    OperationPublicAPI,
			want: "发布接口",
		},
		{
			name: "上线接口",
			o:    OperationUpAPI,
			want: "上线接口",
		},
		{
			name: "下线接口",
			o:    OperationDownAPI,
			want: "下线接口",
		},
		{
			name: "删除接口",
			o:    OperationDeleteAPI,
			want: "删除接口",
		},
		{
			name: "查看接口调用信息",
			o:    OperationGetAPIAuthInfo,
			want: "查看接口调用信息",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.SimplifiedChineseName()
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestFilterOperationsByAuditType(t *testing.T) {
	type args struct {
		in []Operation
		t  AuditType
	}
	tests := []struct {
		name string

		args args
		want []Operation
	}{
		{
			name: "管理日志",
			args: args{
				in: []Operation{
					OperationGenerateAPI,
					OperationRegisterAPI,
					OperationUpdateAPI,
					OperationPublicAPI,
					OperationDeleteAPI,
					OperationCreateDataDownloadTask,
					OperationDeleteDataDownloadTask,
					OperationDataDownload,
					OperationCreateDimensionModel,
					OperationUpdateDimensionModelBasicInfo,
					OperationUpdateDimensionModelConfigInfo,
					OperationDeleteDimensionModel,
					OperationCreateIndicator,
					OperationUpdateIndicator,
					OperationDeleteIndicator,
					OperationQueryIndicatorResult,
				},
				t: AuditTypeManagement,
			},
			want: []Operation{
				OperationGenerateAPI,
				OperationRegisterAPI,
				OperationUpdateAPI,
				OperationPublicAPI,
				OperationDeleteAPI,
				OperationCreateDimensionModel,
				OperationUpdateDimensionModelBasicInfo,
				OperationUpdateDimensionModelConfigInfo,
				OperationDeleteDimensionModel,
				OperationCreateIndicator,
				OperationUpdateIndicator,
				OperationDeleteIndicator,
			},
		},
		{
			name: "操作日志",
			args: args{
				in: []Operation{
					OperationGenerateAPI,
					OperationRegisterAPI,
					OperationUpdateAPI,
					OperationPublicAPI,
					OperationDeleteAPI,
					OperationCreateDataDownloadTask,
					OperationDeleteDataDownloadTask,
					OperationDataDownload,
					OperationCreateDimensionModel,
					OperationUpdateDimensionModelBasicInfo,
					OperationUpdateDimensionModelConfigInfo,
					OperationDeleteDimensionModel,
					OperationCreateIndicator,
					OperationUpdateIndicator,
					OperationDeleteIndicator,
					OperationQueryIndicatorResult,
				},
				t: AuditTypeOperation,
			},
			want: []Operation{
				OperationCreateDataDownloadTask,
				OperationDeleteDataDownloadTask,
				OperationDataDownload,
				OperationQueryIndicatorResult,
			},
		},
		{
			name: "未知类型",
			args: args{
				in: []Operation{
					OperationGenerateAPI,
					OperationRegisterAPI,
					OperationUpdateAPI,
					OperationPublicAPI,
					OperationDeleteAPI,
					OperationCreateDataDownloadTask,
					OperationDeleteDataDownloadTask,
					OperationDataDownload,
					OperationCreateDimensionModel,
					OperationUpdateDimensionModelBasicInfo,
					OperationUpdateDimensionModelConfigInfo,
					OperationDeleteDimensionModel,
					OperationCreateIndicator,
					OperationUpdateIndicator,
					OperationDeleteIndicator,
					OperationQueryIndicatorResult,
				},
				t: "AUDIT_TYPE_UNKNOWN",
			},
		},
		{
			name: "无修改",
			args: args{
				in: []Operation{
					OperationGenerateAPI,
					OperationRegisterAPI,
					OperationUpdateAPI,
					OperationPublicAPI,
					OperationDeleteAPI,
				},
				t: AuditTypeManagement,
			},
			want: []Operation{
				OperationGenerateAPI,
				OperationRegisterAPI,
				OperationUpdateAPI,
				OperationPublicAPI,
				OperationDeleteAPI,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterOperationsByAuditType(tt.args.in, tt.args.t)
			assert.ElementsMatch(t, tt.want, got)
			assert.NotEqual(t, reflect.ValueOf(tt.args.in).Pointer(), reflect.ValueOf(got).Pointer(), "&tt.args.in: %p, &got: %p, 即使无修改也应返回新对象", tt.args.in, got)
		})
	}
}

func TestOperationsForAuditType(t *testing.T) {
	tests := []struct {
		name      string
		auditType AuditType
		want      []Operation
	}{
		{
			name:      "管理日志",
			auditType: AuditTypeManagement,
			want: lo.FilterMap(operationAndAuditTypeBindings, func(b operationAndAuditTypeBinding, _ int) (Operation, bool) {
				return b.operation, b.auditType == AuditTypeManagement
			}),
		},
		{
			name:      "操作日志",
			auditType: AuditTypeOperation,
			want: lo.FilterMap(operationAndAuditTypeBindings, func(b operationAndAuditTypeBinding, _ int) (Operation, bool) {
				return b.operation, b.auditType == AuditTypeOperation
			}),
		},
		{
			name:      "未定义的审计日志类型",
			auditType: "AUDIT_TYPE_UNKNOWN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OperationsForAuditType(tt.auditType)
			assert.Equal(t, tt.want, got)
		})
	}
}
