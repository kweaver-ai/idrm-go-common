package v1

import (
	"net/url"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

// url.Values -> v1.WorkOrderTaskListOptions
func Convert_url_Values_To_v1_WorkOrderTaskListOptions(in *url.Values, out *WorkOrderTaskListOptions) error {
	if err := meta_v1.Convert_url_Values_To_v1_ListOptions(in, &out.ListOptions); err != nil {
		return err
	}
	if values, ok := map[string][]string(*in)["keyword"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.Keyword); err != nil {
			return err
		}
	} else {
		out.Keyword = ""
	}
	if values, ok := map[string][]string(*in)["status"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_v1_WorkOrderTaskStatus(&values, &out.Status); err != nil {
			return err
		}
	} else {
		out.Status = ""
	}
	if values, ok := map[string][]string(*in)["work_order_type"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_v1_WorkOrderType(&values, &out.WorkOrderType); err != nil {
			return err
		}
	} else {
		out.WorkOrderType = ""
	}

	if values, ok := map[string][]string(*in)["work_order_id"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.WorkOrderID); err != nil {
			return err
		}
	} else {
		out.WorkOrderID = ""
	}
	return nil
}

// []string -> v1.WorkOrderTaskStatus
func Convert_Slice_string_To_v1_WorkOrderTaskStatus(in *[]string, out *WorkOrderTaskStatus) error {
	if len(*in) > 0 {
		*out = WorkOrderTaskStatus((*in)[0])
	}
	return nil
}

// []string -> v1.WorkOrderType
func Convert_Slice_string_To_v1_WorkOrderType(in *[]string, out *WorkOrderType) error {
	if len(*in) > 0 {
		*out = WorkOrderType((*in)[0])
	}
	return nil
}
