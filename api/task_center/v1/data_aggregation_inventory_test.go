package v1

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func TestDataAggregationInventory_MarshalJSON(t *testing.T) {
	inventory := &DataAggregationInventory{
		Name:         "测试：新建 - 暂存",
		DepartmentID: "29d7968c-d308-11ef-bc0e-46740f10a1c5",
		Resources: []DataAggregationResource{
			{
				DataViewID:         "82781cac-a345-499b-8e95-d9beceee3a45",
				CollectionMethod:   DataAggregationResourceCollectionFull,
				SyncFrequency:      DataAggregationResourceSyncFrequencyPerDay,
				TargetDatasourceID: "2483ed28-f4ba-4026-8049-c20ec8a895f3",
			},
		},
	}

	inventoryJSON, err := json.MarshalIndent(inventory, "", "  ")
	require.NoError(t, err)
	t.Logf("DataAggregationInventory: %s", inventoryJSON)
}

func TestDataAggregationInventoryListOptions_UnmarshalQuery(t *testing.T) {
	tests := []struct {
		name  string
		query string
		want  *DataAggregationInventoryListOptions
	}{
		{
			name:  "example",
			query: "limit=1453&offset=1644&sort=name&direction=desc&keyword=hello&fields=code,name&status=Auditing,Completed,Draft&department_ids=id_0,id_1",
			want: &DataAggregationInventoryListOptions{
				ListOptions: meta_v1.ListOptions{
					Offset:    1644,
					Limit:     1453,
					Sort:      "name",
					Direction: meta_v1.Descending,
				},
				Keyword: "hello",
				Fields: []DataAggregationInventoryListKeywordField{
					DataAggregationInventoryListKeywordFieldCode,
					DataAggregationInventoryListKeywordFieldName,
				},
				Statuses: []DataAggregationInventoryStatus{
					DataAggregationInventoryAuditing,
					DataAggregationInventoryCompleted,
					DataAggregationInventoryDraft,
				},
				DepartmentIDs: []string{
					"id_0",
					"id_1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := url.ParseQuery(tt.query)
			require.NoError(t, err)

			opts := &DataAggregationInventoryListOptions{}
			require.NoError(t, opts.UnmarshalQuery(data))

			assert.Equal(t, tt.want, opts)
		})
	}
}
