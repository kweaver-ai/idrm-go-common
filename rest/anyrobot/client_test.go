package anyrobot

import (
	"context"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	dataModelV1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/data-model/v1"
	v1 "github.com/kweaver-ai/idrm-go-common/api/anyrobot/uniquery/v1"
)

func TestClient(t *testing.T) {
	var ctx = context.Background()

	// anyrobot server address
	server := test.ShouldGetenv(t, "TEST_SERVER")
	base, err := url.Parse(server)
	if err != nil {
		t.Fatal(err)
	}
	// http client
	client := http.DefaultClient

	cs := New(base, client)

	resp, err := cs.UniqueryV1().DataViews().Get(
		ctx,
		[]string{"514940169606948879"},
		&v1.DataViewQuery{
			Filters: &v1.DataViewQueryFilters{
				Operation: "like",
				Field:     "Body.resource_operation.operator.name",
				Value:     "Scarlet",
				ValueFrom: v1.ValueFromConst,
			},
			Offset: 1,
			Limit:  2,
			Format: v1.DataViewQueryOriginal,
		},
		&v1.DataViewQueryOptions{},
	)
	if err != nil {
		t.Fatal(err)
	}

	if !assert.Equal(t, 1, len(resp.Datas)) {
		return
	}

	// datas[0]
	var data = resp.Datas[0]

	assert.Equal(t, "12", data.Total, "datas[0].total")
	assert.NotEmpty(t, data.Values, "datas[0].values")

	// datas[0].values[0]
	var value = data.Values[0]

	assert.Equal(t, time.Date(2024, 8, 23, 9, 26, 6, 788000000, time.UTC).Truncate(time.Millisecond), value.Timestamp, "datas[0].values[0].@timestamp")
	assert.Equal(t, "Info", value.SeverityText, "datas[0].values[0].SeverityText")
}

func TestDataModelV1DataViewGet(t *testing.T) {
	// testdata
	var (
		server = test.ShouldGetenv(t, "TEST_SERVER")

		dataViewID = test.ShouldGetenv(t, "TEST_DATA_VIEW_ID")

		want = []dataModelV1.DataView{
			{
				ID: "537245430564585106",
				Fields: []dataModelV1.Field{
					{Name: "__data_type"},
					{Name: "__index_base"},
					{Name: "__write_time"},
					{Name: "@timestamp"},
					{Name: "@version"},
					{Name: "Attributes"},
					{Name: "Body.FieldType"},
					{Name: "Body.resource_operation.description"},
					{Name: "Body.resource_operation.operation"},
					{Name: "Body.resource_operation.operator.agent.ip"},
					{Name: "Body.resource_operation.operator.agent.ips"},
					{Name: "Body.resource_operation.operator.agent.type"},
					{Name: "Body.resource_operation.operator.department.id"},
					{Name: "Body.resource_operation.operator.department.name"},
					{Name: "Body.resource_operation.operator.id"},
					{Name: "Body.resource_operation.operator.name"},
					{Name: "Body.resource_operation.operator.type"},
					{Name: "Body.resource_operation.source_object.department_path"},
					{Name: "Body.resource_operation.source_object.id"},
					{Name: "Body.resource_operation.source_object.name"},
					{Name: "Body.resource_operation.source_object.owner_name"},
					{Name: "Body.resource_operation.target_object.department_path"},
					{Name: "Body.resource_operation.target_object.id"},
					{Name: "Body.resource_operation.target_object.name"},
					{Name: "Body.resource_operation.target_object.owner_name"},
					{Name: "Body.resource_operation.target_object.subject_path"},
					{Name: "category"},
					{Name: "geoip.ip"},
					{Name: "geoip.latitude"},
					{Name: "geoip.location"},
					{Name: "geoip.longitude"},
					{Name: "Link.SpanId"},
					{Name: "Link.TraceId"},
					{Name: "Resource.job_id"},
					{Name: "Resource.service.instance.id"},
					{Name: "Resource.service.name"},
					{Name: "Resource.service.version"},
					{Name: "Resource.telemetry.sdk.language"},
					{Name: "Resource.telemetry.sdk.name"},
					{Name: "Resource.telemetry.sdk.version"},
					{Name: "SeverityText"},
					{Name: "tags"},
					{Name: "type"},
				},
			},
		}
	)

	base, err := url.Parse(server)
	if err != nil {
		t.Fatal(err)
	}

	var cs Interface = New(base, http.DefaultClient)

	got, err := cs.DataModelV1().DataViews().Get(context.Background(), []string{dataViewID})
	assert.NoError(t, err)
	if assert.Len(t, got, 1) {
		assert.Equal(t, dataViewID, got[0].ID, "[0].ID")
		assert.ElementsMatch(t, want[0].Fields, got[0].Fields, "[0].Fields")
	}
}
