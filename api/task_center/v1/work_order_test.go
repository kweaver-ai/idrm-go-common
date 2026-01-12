package v1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkOrderUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		want *WorkOrder
	}{
		{
			name: "create.standalone",
			want: &WorkOrder{
				Name:                     "无 0512 1713",
				Type:                     WorkOrderDataAggregation,
				ResponsibleUID:           "8ca13f64-ff06-11ef-bbe0-ba020467f4e3",
				Priority:                 WorkOrderCommon,
				Description:              "说明",
				Remark:                   "备注",
				SourceType:               WorkOrderStandalone,
				DataAggregationInventory: &DataAggregationInventory{ID: "0196c329-661f-7413-bf8f-b8b3e6bbecd5"},
			},
		},
		{
			name: "create.plan",
			want: &WorkOrder{
				Name:                     "计划 0514 1624",
				Type:                     WorkOrderDataAggregation,
				ResponsibleUID:           "8ca13f64-ff06-11ef-bbe0-ba020467f4e3",
				Priority:                 WorkOrderCommon,
				Description:              "说明",
				Remark:                   "备注",
				SourceType:               WorkOrderPlan,
				SourceID:                 "cdb59a06-44ad-4067-844a-c21aec81db73",
				DataAggregationInventory: &DataAggregationInventory{ID: "0196c329-661f-7413-bf8f-b8b3e6bbecd5"},
			},
		},
		{
			name: "create.business_form",
			want: &WorkOrder{
				Name:           "业务表 0512 1826",
				Type:           WorkOrderDataAggregation,
				ResponsibleUID: "8ca13f64-ff06-11ef-bbe0-ba020467f4e3",
				Priority:       WorkOrderCommon,
				Description:    "说明",
				Remark:         "备注",
				SourceType:     WorkOrderBusinessForm,
				SourceIDs: []string{
					"3a9d7fcf-9b24-401b-acaf-525e5eeae097",
					"7511f4a2-f8d3-419a-9893-908ce5d1b4d9",
					"832b4ead-ae8c-4ec5-b6c8-ac31f7eb1d46",
				},
				DataAggregationInventory: &DataAggregationInventory{
					Resources: []DataAggregationResource{
						{
							DataViewID:         "fdf3d53c-a829-4e78-b6b9-51fedc5ffd5e",
							CollectionMethod:   DataAggregationResourceCollectionFull,
							SyncFrequency:      DataAggregationResourceSyncFrequencyPerHour,
							BusinessFormID:     "3a9d7fcf-9b24-401b-acaf-525e5eeae097",
							TargetDatasourceID: "e2531535-629a-42f5-9d5c-c8a7fc2a74e5",
						},
						{
							DataViewID:         "af2ea0a4-7f37-4588-8ba3-2afac40e1d3c",
							CollectionMethod:   DataAggregationResourceCollectionFull,
							SyncFrequency:      DataAggregationResourceSyncFrequencyPerHour,
							BusinessFormID:     "7511f4a2-f8d3-419a-9893-908ce5d1b4d9",
							TargetDatasourceID: "e2531535-629a-42f5-9d5c-c8a7fc2a74e5",
						},
						{
							DataViewID:         "a68e78f4-130a-408b-96de-a5def43383b7",
							CollectionMethod:   DataAggregationResourceCollectionFull,
							SyncFrequency:      DataAggregationResourceSyncFrequencyPerHour,
							BusinessFormID:     "832b4ead-ae8c-4ec5-b6c8-ac31f7eb1d46",
							TargetDatasourceID: "e2531535-629a-42f5-9d5c-c8a7fc2a74e5",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", "work_order."+tt.name+".json"))
			require.NoError(t, err)
			defer f.Close()

			got := &WorkOrder{}
			require.NoError(t, json.NewDecoder(f).Decode(got))

			assert.Equal(t, tt.want, got)
		})
	}
}
