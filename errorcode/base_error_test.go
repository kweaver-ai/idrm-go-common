package errorcode

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func TestNewErrorForHTTPResponse(t *testing.T) {
	tests := []struct {
		name string
		err  *meta_v1.Error
	}{
		{
			name: "example",
			err: &meta_v1.Error{
				Code:        "DemandManagement.Public.InvalidParameter",
				Description: "参数值校验不通过",
				Solution:    "请使用请求参数构造规范化的请求字符串。详细信息参见产品 API 文档",
				Detail:      json.RawMessage(`[{"key":"catalog_ids","message":"catalog_ids为必填字段"}]`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errJSON, err := json.Marshal(tt.err)
			require.NoError(t, err)

			resp := &http.Response{Body: io.NopCloser(bytes.NewReader(errJSON))}
			got := NewErrorForHTTPResponse(resp)
			assert.Equal(t, tt.err, got)
		})
	}
}
