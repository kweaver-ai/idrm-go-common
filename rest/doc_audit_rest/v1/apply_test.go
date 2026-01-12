package v1

import (
	"crypto/tls"
	"net/http"
	"net/http/httputil"
	"testing"
)

type debuggingRoundTripper struct {
	t *testing.T
}

// RoundTrip implements http.RoundTripper.
func (rt *debuggingRoundTripper) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	rt.t.Helper()

	if b, err := httputil.DumpRequestOut(req, true); err != nil {
		rt.t.Errorf("dump http request fail: %v", err)
	} else {
		rt.t.Logf("http request:\n%s", b)
	}

	underlying := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	resp, err = underlying.RoundTrip(req)
	if err != nil {
		return
	}

	if b, err := httputil.DumpResponse(resp, true); err != nil {
		rt.t.Errorf("dump http response fail: %v", err)
	} else {
		rt.t.Logf("http response:\n%s", b)
	}

	return
}

var _ http.RoundTripper = &debuggingRoundTripper{}
