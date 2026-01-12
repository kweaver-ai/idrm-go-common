package v1

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func TestIndicatorDimensionalRuleSpec_MarshalJSON(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "indicator_dimensional_rule_spec.json"))
	require.NoError(t, err)

	spec := &IndicatorDimensionalRuleSpec{
		IndicatorID: 12450,
		Rule: Rule{
			Name: "NAME",
			Fields: []Field{
				{
					ID:       "00000000-0000-0000-0000-111111111111",
					Name:     "名称 0000",
					NameEn:   "NAME EN 0000",
					DataType: "DATA TYPE 0000",
				},
				{
					ID:       "00000000-0000-0000-1111-111111111111",
					Name:     "名称 1111",
					NameEn:   "NAME EN 1111",
					DataType: "DATA TYPE 1111",
				},
			},
		},
	}
	got, err := json.Marshal(spec)
	require.NoError(t, err)

	assert.JSONEq(t, string(want), string(got))
}

func TestIndicatorDimensionalRuleSpec_UnmarshalJSON(t *testing.T) {
	want := &IndicatorDimensionalRuleSpec{
		IndicatorID: 12450,
		Rule: Rule{
			Name: "NAME",
			Fields: []Field{
				{
					ID:       "00000000-0000-0000-0000-111111111111",
					Name:     "名称 0000",
					NameEn:   "NAME EN 0000",
					DataType: "DATA TYPE 0000",
				},
				{
					ID:       "00000000-0000-0000-1111-111111111111",
					Name:     "名称 1111",
					NameEn:   "NAME EN 1111",
					DataType: "DATA TYPE 1111",
				},
			},
		},
	}

	data, err := os.ReadFile(filepath.Join("testdata", "indicator_dimensional_rule_spec.json"))
	require.NoError(t, err)

	got := &IndicatorDimensionalRuleSpec{}
	require.NoError(t, json.Unmarshal(data, got))

	assert.Equal(t, want, got)
}

func TestIndicatorDimensionalRuleListOptions_UnmarshalQuery(t *testing.T) {
	tests := []struct {
		name    string
		data    url.Values
		want    IndicatorDimensionalRuleListOptions
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "unspecified",
			want:    IndicatorDimensionalRuleListOptions{},
			wantErr: require.NoError,
		},
		{
			name:    "specified indicator id",
			data:    url.Values{"indicator_id": []string{"1111"}},
			want:    IndicatorDimensionalRuleListOptions{IndicatorID: 1111},
			wantErr: require.NoError,
		},
		{
			name: "specified multi indicator ids",
			data: url.Values{"indicator_id": []string{
				"1111",
				"2222",
				"3333",
			}},
			want:    IndicatorDimensionalRuleListOptions{IndicatorID: 3333},
			wantErr: require.NoError,
		},
		{
			name:    "invalid indicator id",
			data:    url.Values{"indicator_id": []string{"xxxx"}},
			wantErr: require.Error,
		},
		{
			name: "sort",
			data: url.Values{"sort": []string{"is_authorized"}, "direction": []string{"desc"}},
			want: IndicatorDimensionalRuleListOptions{
				ListOptions: meta_v1.ListOptions{
					Sort:      string(IndicatorDimensionalRuleSortIsAuthorized),
					Direction: meta_v1.Descending,
				},
				Sort:      IndicatorDimensionalRuleSortIsAuthorized,
				Direction: meta_v1.Descending,
			},
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts IndicatorDimensionalRuleListOptions
			tt.wantErr(t, opts.UnmarshalQuery(tt.data))
			assert.Equal(t, tt.want, opts)
		})
	}
}
