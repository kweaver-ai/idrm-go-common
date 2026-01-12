package v1

import (
	"net/url"
	"strconv"
	"strings"

	v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	meta "github.com/kweaver-ai/idrm-go-common/api/meta/v1"
)

func queryForListOptions(opts *meta.ListOptions) (q url.Values) {
	q = make(url.Values)

	if opts.Offset > 0 {
		q.Add("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		q.Add("limit", strconv.Itoa(opts.Limit))
	}

	return
}

func queryForApplyListOptions(opts *v1.ApplyListOptions) (q url.Values) {
	q = queryForListOptions(&opts.ListOptions)

	if len(opts.Type) > 0 {
		var typeStrings []string
		for _, t := range opts.Type {
			typeStrings = append(typeStrings, string(t))
		}
		q.Add("type", strings.Join(typeStrings, ","))
	}

	if opts.Status != "" {
		q.Add("status", string(opts.Status))
	}

	return
}

func queryForHistoryListOptions(opts *v1.HistoryListOptions) (q url.Values) {
	q = queryForListOptions(&opts.ListOptions)

	if len(opts.Type) > 0 {
		var typeStrings []string
		for _, t := range opts.Type {
			typeStrings = append(typeStrings, string(t))
		}
		q.Add("type", strings.Join(typeStrings, ","))
	}

	if opts.Status != "" {
		q.Add("status", string(opts.Status))
	}

	return
}

func queryForTaskListOptions(opts *v1.TaskListOptions) (q url.Values) {
	q = queryForListOptions(&opts.ListOptions)

	if len(opts.Type) > 0 {
		var typeStrings []string
		for _, t := range opts.Type {
			typeStrings = append(typeStrings, string(t))
		}
		q.Add("type", strings.Join(typeStrings, ","))
	}

	if len(opts.Status) > 0 {
		q.Add("status", string(opts.Status))
	}

	return
}
