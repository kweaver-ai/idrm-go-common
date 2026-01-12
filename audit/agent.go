package audit

import (
	"net"
	"net/http"
	"strings"

	"github.com/mileusna/useragent"

	v1 "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
)

// AgentFromRequest returns an Agent based on specific http request.
func AgentFromRequest(req *http.Request) v1.Agent {
	return v1.Agent{
		Type: agentTypeFromUserAgent(req.UserAgent()),
		IP:   sourceIP(req),
	}
}

// agentTypeFromUserAgent returns agent type based on User-Agent.
//
// If UserAgent belongs to agentWebUserAgentNameSet or contains the string
// "browser" (case insensitive) it is considered AgentWeb.
func agentTypeFromUserAgent(userAgent string) v1.AgentType {
	ua := useragent.Parse(userAgent)
	switch {
	case agentWebUserAgentNameSet.Has(ua.Name):
		return v1.AgentWeb
	case strings.Contains(strings.ToLower(ua.Name), "browser"):
		return v1.AgentWeb
	default:
		return v1.AgentUnknown
	}
}

const headerKeyXForwardedFor = "X-Forwarded-For"

func sourceIP(req *http.Request) (ip net.IP) {
	for _, value := range req.Header.Values(headerKeyXForwardedFor) {
		for _, v := range strings.Split(value, ",") {
			if v = strings.TrimSpace(v); v == "" {
				continue
			}

			if ip = net.ParseIP(v); ip == nil {
				continue
			}

			return
		}
	}

	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	ip = net.ParseIP(host)
	return
}
