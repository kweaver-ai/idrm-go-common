package v1

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalDataView(t *testing.T) {
	data, err := os.ReadFile(filepath.Join("testdata", "data-views", "537245430564585106.json"))
	if err != nil {
		t.Fatal(err)
	}

	got := &DataView{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatal(err)
	}

	want := &DataView{
		ID: "537245430564585106",
		Fields: []Field{
			{Name: "@timestamp"},
			{Name: "__data_type"},
			{Name: "__index_base"},
			{Name: "__write_time"},
			{Name: "category"},
			{Name: "tags"},
			{Name: "type"},
			{Name: "Resource.job_id"},
			{Name: "Resource.service.name"},
			{Name: "Resource.service.version"},
			{Name: "Resource.service.instance.id"},
			{Name: "Resource.telemetry.sdk.version"},
			{Name: "Resource.telemetry.sdk.language"},
			{Name: "Resource.telemetry.sdk.name"},
			{Name: "Link.TraceId"},
			{Name: "Link.SpanId"},
			{Name: "SeverityText"},
			{Name: "@version"},
			{Name: "Body.FieldType"},
			{Name: "Body.resource_operation.description"},
			{Name: "Body.resource_operation.operation"},
			{Name: "Body.resource_operation.operator.department.id"},
			{Name: "Body.resource_operation.operator.department.name"},
			{Name: "Body.resource_operation.operator.id"},
			{Name: "Body.resource_operation.operator.name"},
			{Name: "Body.resource_operation.operator.type"},
			{Name: "Body.resource_operation.operator.agent.type"},
			{Name: "Body.resource_operation.operator.agent.ip"},
			{Name: "Body.resource_operation.operator.agent.ips"},
			{Name: "Body.resource_operation.source_object.name"},
			{Name: "Body.resource_operation.source_object.id"},
			{Name: "Body.resource_operation.source_object.owner_name"},
			{Name: "Body.resource_operation.source_object.department_path"},
			{Name: "Body.resource_operation.target_object.department_path"},
			{Name: "Body.resource_operation.target_object.id"},
			{Name: "Body.resource_operation.target_object.name"},
			{Name: "Body.resource_operation.target_object.owner_name"},
			{Name: "geoip.location"},
			{Name: "geoip.longitude"},
			{Name: "geoip.ip"},
			{Name: "geoip.latitude"},
			{Name: "Attributes"},
		},
	}
	assert.Equal(t, want, got)
}
