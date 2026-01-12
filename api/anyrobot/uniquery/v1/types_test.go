package v1

import (
	_ "embed"
	"encoding/json"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/resp.json
var respJSONBytes []byte

func TestViewUniResponse(t *testing.T) {
	resp := &ViewUniResponse{}
	if err := json.Unmarshal(respJSONBytes, resp); err != nil {
		t.Fatal(err)
	}

	if !assert.Equal(t, 1, len(resp.Datas)) {
		return
	}

	// datas[0]
	var data = resp.Datas[0]

	assert.Equal(t, "8", data.Total, "datas[0].total")
	assert.NotEmpty(t, data.Values, "datas[0].values")

	// datas[0].values[0]
	var value = data.Values[0]

	assert.Equal(t, time.Date(2024, 07, 05, 01, 45, 49, 293000000, time.UTC).Truncate(time.Millisecond), value.Timestamp, "datas[0].values[0].@timestamp")
	assert.Equal(t, "Info", value.SeverityText, "datas[0].values[0].SeverityText")
}

func TestUnmarshalStatus(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     *Status
	}{
		{
			name:     "过滤条件个数超出限制",
			filename: "uniquery_data_view_count_exceeded_filters.json",
			want: &Status{
				ErrorCode:    "Uniquery.DataView.CountExceeded.Filters",
				Description:  "过滤条件个数超出限制",
				Solution:     "请检查参数是否正确。",
				ErrorLink:    "暂无",
				ErrorDetails: "The number of subConditions exceeds 10",
			},
		},
		{
			name:     "过滤器配置无效",
			filename: "uniquery_data_view_invalid_parameter_filters.json",
			want: &Status{
				ErrorCode:    "Uniquery.DataView.InvalidParameter.Filters",
				Description:  "过滤器配置无效",
				Solution:     "请检查参数是否正确。",
				ErrorLink:    "暂无",
				ErrorDetails: "failed to new condition, condition config field name must in view original fields",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(path.Join("testdata", "errors", tt.filename))
			if err != nil {
				t.Fatal(err)
			}

			got := &Status{}
			assert.NoError(t, json.Unmarshal(data, got))
			assert.Equal(t, tt.want, got)
		})
	}
}
