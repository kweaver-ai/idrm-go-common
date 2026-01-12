package v1

import (
	"net/url"
	"strings"

	meta_v1 "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func Convert_V1_RoleListOptions_To_url_Values(in *RoleListOptions, out *url.Values) error {
	if err := meta_v1.Convert_v1_ListOptions_To_url_Values(&in.ListOptions, out); err != nil {
		return err
	}
	if in.Keyword != "" {
		out.Add("keyword", in.Keyword)
	}
	if in.Type != "" {
		out.Add("type", string(in.Type))
	}
	if in.RoleGroupID != "" {
		out.Add("role_group_id", in.RoleGroupID)
	}
	if len(in.UserIDs) != 0 {
		out.Add("user_ids", strings.Join(in.UserIDs, ","))
	}
	return nil
}

func Convert_V1_RoleGroupListOptions_To_url_Values(in *RoleGroupListOptions, out *url.Values) error {
	if err := meta_v1.Convert_v1_ListOptions_To_url_Values(&in.ListOptions, out); err != nil {
		return err
	}
	if in.Keyword != "" {
		out.Add("keyword", in.Keyword)
	}
	if len(in.UserIDs) != 0 {
		out.Add("user_ids", strings.Join(in.UserIDs, ","))
	}
	return nil
}

func Convert_V1_UserListOptions_To_url_Values(in *UserListOptions, out *url.Values) error {
	if err := meta_v1.Convert_v1_ListOptions_To_url_Values(&in.ListOptions, out); err != nil {
		return err
	}
	if in.Keyword != "" {
		out.Add("keyword", in.Keyword)
	}
	if in.DepartmentID != "" {
		out.Add("department_id", in.DepartmentID)
	}
	return nil
}

func Convert_url_Values_To_V1_RoleListOptions(in *url.Values, out *RoleListOptions) error {
	if err := meta_v1.Convert_url_Values_To_v1_ListOptions(in, &out.ListOptions); err != nil {
		return err
	}
	if values, ok := (*in)["keyword"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.Keyword); err != nil {
			return err
		}
	} else {
		out.Keyword = ""
	}
	if values, ok := (*in)["type"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_generic_string(&values, &out.Type); err != nil {
			return err
		}
	} else {
		out.Type = ""
	}
	if values, ok := (*in)["role_group_id"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.RoleGroupID); err != nil {
			return err
		}
	} else {
		out.RoleGroupID = ""
	}
	for _, v := range (*in)["user_ids"] {
		out.UserIDs = append(out.UserIDs, strings.Split(v, ",")...)
	}
	return nil
}

func Convert_url_Values_To_V1_RoleGroupListOptions(in *url.Values, out *RoleGroupListOptions) error {
	if err := meta_v1.Convert_url_Values_To_v1_ListOptions(in, &out.ListOptions); err != nil {
		return err
	}
	if values, ok := (*in)["keyword"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.Keyword); err != nil {
			return err
		}
	} else {
		out.Keyword = ""
	}
	for _, v := range (*in)["user_ids"] {
		out.UserIDs = append(out.UserIDs, strings.Split(v, ",")...)
	}
	return nil
}

func Convert_url_Values_To_V1_UserListOptions(in *url.Values, out *UserListOptions) error {
	if err := meta_v1.Convert_url_Values_To_v1_ListOptions(in, &out.ListOptions); err != nil {
		return err
	}
	if values, ok := (*in)["keyword"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.Keyword); err != nil {
			return err
		}
	} else {
		out.Keyword = ""
	}
	if values, ok := (*in)["department_id"]; ok && len(values) > 0 {
		if err := meta_v1.Convert_Slice_string_To_string(&values, &out.DepartmentID); err != nil {
			return err
		}
	} else {
		out.DepartmentID = ""
	}
	return nil
}
