package v1

import (
	"net/url"
	"strconv"
	"strings"
	_ "strings"
)

// url.Values -> ListOptions
func Convert_url_Values_To_v1_ListOptions(in *url.Values, out *ListOptions) error {
	if values, ok := map[string][]string(*in)["offset"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_int(&values, &out.Offset); err != nil {
			return err
		}
	} else {
		out.Offset = 0
	}

	if values, ok := map[string][]string(*in)["limit"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_int(&values, &out.Limit); err != nil {
			return err
		}
	} else {
		out.Limit = 0
	}
	if values, ok := map[string][]string(*in)["sort"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_string(&values, &out.Sort); err != nil {
			return err
		}
	} else {
		out.Sort = ""
	}
	if values, ok := map[string][]string(*in)["direction"]; ok && len(values) > 0 {
		if err := Convert_Slice_string_To_v1_Direction(&values, &out.Direction); err != nil {
			return err
		}
	} else {
		out.Direction = ""
	}
	return nil
}

// []string -> v1.Direction
func Convert_Slice_string_To_v1_Direction(in *[]string, out *Direction) error {
	if len(*in) > 0 {
		*out = Direction((*in)[0])
	}
	return nil
}

// []string -> string
func Convert_Slice_string_To_string(in *[]string, out *string) error {
	if len(*in) == 0 {
		*out = ""
		return nil
	}
	*out = (*in)[0]
	return nil
}

// []string -> int
func Convert_Slice_string_To_int(in *[]string, out *int) error {
	if len(*in) == 0 {
		*out = 0
		return nil
	}
	str := (*in)[0]
	i, err := strconv.Atoi(str)
	if err != nil {
		return err
	}
	*out = i
	return nil
}

func Convert_v1_ListOptions_To_url_Values(in *ListOptions, out *url.Values) error {
	if in.Offset != 0 {
		out.Add("offset", strconv.Itoa(in.Offset))
	}
	if in.Limit != 0 {
		out.Add("limit", strconv.Itoa(in.Limit))
	}
	if in.Sort != "" {
		out.Add("sort", in.Sort)
	}
	if in.Direction != "" {
		out.Add("direction", string(in.Direction))
	}
	return nil
}

func Convert_Slice_string_To_generic_string[T ~string](in *[]string, out *T) error {
	if len(*in) == 0 {
		*out = ""
		return nil
	}
	*out = T((*in)[0])
	return nil
}

// Convert_Slice_string_To_bool will convert a string parameter to boolean.
// Only the absence of a value (i.e. zero-length slice), a value of "false", or a
// value of "0" resolve to false.
// Any other value (including empty string) resolves to true.
func Convert_Slice_string_To_bool(in *[]string, out *bool) error {
	if len(*in) == 0 {
		*out = false
		return nil
	}
	switch {
	case (*in)[0] == "0", strings.EqualFold((*in)[0], "false"):
		*out = false
	default:
		*out = true
	}
	return nil
}
