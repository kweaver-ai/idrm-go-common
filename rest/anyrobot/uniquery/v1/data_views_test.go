package v1

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	v1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/uniquery/v1"
	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

type httpRequestAssertionRoundTripper struct {
	// want http method
	method string
	// want http url
	url *url.URL
	// want http body
	body []byte

	assertions *assert.Assertions
	// underlying
	rt http.RoundTripper
}

// RoundTrip implements http.RoundTripper.
func (rt *httpRequestAssertionRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// cache request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("cache request body fail: %w", err)
	}
	req.Body.Close()
	req.Body = io.NopCloser(bytes.NewReader(body))

	// assert method
	rt.assertions.Equal(rt.method, req.Method, "Method")
	// assert url
	rt.assertions.Equal(rt.url, req.URL, "URL")
	// assert body
	rt.assertions.JSONEq(string(rt.body), string(body), "Body")

	resp, err = rt.rt.RoundTrip(req)
	return
}

var _ http.RoundTripper = &httpRequestAssertionRoundTripper{}

type httpResponseRoundTripper struct {
	response *http.Response
	err      error
}

// RoundTrip implements http.RoundTripper.
func (rt *httpResponseRoundTripper) RoundTrip(*http.Request) (resp *http.Response, err error) {
	return rt.response, rt.err
}

var _ http.RoundTripper = &httpResponseRoundTripper{}

var (
	//go:embed testdata/request.json
	requestJSON []byte
	//go:embed testdata/response.json
	responseJSON []byte
)

func TestDataViews_Get(t *testing.T) {
	// testdata
	const (
		anyrobotHost = "https://testing.example.org/anyrobot"
		dataViewID   = "514940169606948879"
	)
	var (
		ctx = context.Background()

		time_20240705_094000 = time.Date(2024, 7, 5, 9, 40, 0, 0, time.Local)
		time_20240705_095000 = time.Date(2024, 7, 5, 9, 50, 0, 0, time.Local)
	)

	var rt http.RoundTripper
	rt = &httpResponseRoundTripper{
		response: &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(responseJSON)),
		},
	}
	// 验证 DataViewClient 发送符合期望的 HTTP Request
	rt = &httpRequestAssertionRoundTripper{
		method: http.MethodPost,
		url: &url.URL{
			Scheme: "https",
			Host:   "testing.example.org",
			Path:   "/anyrobot/api/uniquery/v1/data-views/514940169606948879",
		},
		body:       requestJSON,
		assertions: assert.New(t),
		rt:         rt,
	}

	c := &DataViewClient{
		base: &url.URL{
			Scheme: "https",
			Host:   "testing.example.org",
			Path:   "/anyrobot/api/uniquery/v1/data-views",
		},
		client: &http.Client{
			Transport: &httpRequestAssertionRoundTripper{
				method: http.MethodPost,
				url: &url.URL{
					Scheme: "https",
					Host:   "testing.example.org",
					Path:   "/anyrobot/api/uniquery/v1/data-views/514940169606948879",
				},
				body:       requestJSON,
				assertions: assert.New(t),
				rt:         rt,
			},
		},
	}

	q := &v1.DataViewQuery{
		Filters: &v1.DataViewQueryFilters{
			Operation: "==",
			Field:     "Body.FieldType",
			Value:     "resource_operation",
			ValueFrom: v1.ValueFromConst,
		},
		Start:  &meta_v1.TimestampUnixMilli{Time: time_20240705_094000},
		End:    &meta_v1.TimestampUnixMilli{Time: time_20240705_095000},
		Format: v1.DataViewQueryOriginal,
		Offset: 20,
		Limit:  20,
	}

	r, err := c.Get(ctx, []string{dataViewID}, q, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, r.Datas[0].Total, "15", "datas[0].total")
	assert.Equal(t, len(r.Datas[0].Values), 10, "datas[0].total | length")
}

func Test_newQueryForOptions(t *testing.T) {
	tests := []struct {
		name string
		opts *v1.DataViewQueryOptions
		want url.Values
	}{
		{
			name: "nil",
			opts: nil,
			want: nil,
		},
		{
			name: "empty",
			opts: &v1.DataViewQueryOptions{},
			want: nil,
		},
		{
			name: "AllowNonExistField",
			opts: &v1.DataViewQueryOptions{
				AllowNonExistField: true,
			},
			want: url.Values{
				"allow_non_exist_field": []string{"true"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newQueryForOptions(tt.opts)
			assert.Equal(t, tt.want, got)
		})
	}
}
