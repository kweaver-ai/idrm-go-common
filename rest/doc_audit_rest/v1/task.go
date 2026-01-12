package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	v1 "github.com/kweaver-ai/idrm-go-common/api/doc_audit_rest/v1"
	"github.com/kweaver-ai/idrm-go-common/interception"
)

type TasksGetter interface {
	Tasks() TaskInterface
}

type TaskInterface interface {
	List(ctx context.Context, opts *v1.TaskListOptions) (*v1.TaskList, error)
}

type task struct {
	base   *url.URL
	client *http.Client
}

func newTask(c *docAudit) *task {
	b := *c.base
	b.Path = path.Join(b.Path, "tasks")
	return &task{base: &b, client: c.client}
}

// List implements TaskInterface.
func (c *task) List(ctx context.Context, opts *v1.TaskListOptions) (*v1.TaskList, error) {
	// create http request
	b := *c.base
	if opts != nil {
		b.RawQuery = queryForTaskListOptions(opts).Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, b.String(), http.NoBody)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}
	if t, err := interception.BearerTokenFromContextCompatible(ctx); err == nil {
		req.Header.Add("Authorization", "Bearer "+t)
	}

	// send http request
	resp, err := c.client.Do(req)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}
	defer resp.Body.Close()

	// process http response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: Define errorcode
		return nil, err
	}

	// parse http response
	result := &v1.TaskList{}
	if err := json.Unmarshal(body, result); err != nil {
		// TODO: Define errorcode
		return nil, err
	}

	return result, nil
}
