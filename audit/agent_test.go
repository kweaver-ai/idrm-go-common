package audit

import (
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	audit "github.com/kweaver-ai/idrm-go-common/api/audit/v1"
)

func TestAgentFromRequest(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want audit.Agent
	}{
		{
			name: "example",
			req: &http.Request{
				Header: map[string][]string{
					"X-Forwarded-For": {"10.4.37.39, 10.4.134.16"},
					"X-Real-Ip":       {"10.4.37.39"},
					"User-Agent":      {"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:123.0) Gecko/20100101 Firefox/123.0"},
				},
				RemoteAddr: "10.0.2.100:32968",
			},
			want: audit.Agent{
				Type: audit.AgentWeb,
				IP:   net.ParseIP("10.4.37.39"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AgentFromRequest(tt.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_agentTypeFromUserAgent(t *testing.T) {
	tests := []struct {
		name      string
		userAgent string
		want      audit.AgentType
	}{

		{
			userAgent: "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Opera/9.80 (Macintosh; Intel Mac OS X; U; en) Presto/2.2.15 Version/10.00",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Opera/9.60 (Windows NT 6.0; U; en) Presto/2.1.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36 Edg/91.0.864.59",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 13_5_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Mobile/15E148 Safari/604.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			want:      audit.AgentUnknown,
		},
		{
			userAgent: "Mozilla/5.0 (compatible; YandexAccessibilityBot/3.0; +http://yandex.com/bots)",
			want:      audit.AgentUnknown,
		},
		{
			userAgent: "curl/7.64.1",
			want:      audit.AgentUnknown,
		},
		{
			userAgent: "PostmanRuntime/7.26.5",
			want:      audit.AgentUnknown,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36 OPR/46.0.2597.57",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.91 Safari/537.36 Vivaldi/1.92.917.39",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2; GWX:RED)",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322) NS8/0.9.6",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/60.0.3112.89 Mobile/14F89 Safari/602.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 9_3 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/14.0.0.104835 Mobile/13E233 Safari/9537.53",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/58.0.3029.113 Mobile/14F89 Safari/602.1",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.91 Mobile Safari/537.36 OPR/42.9.2246.119956",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Opera/9.80 (Android; Opera Mini/28.0.2254/66.318; U; en) Presto/2.12.423 Version/12.16",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Linux; U; Android 4.3; en-us; GT-I9300 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			want:      audit.AgentWeb,
		},
		{
			userAgent: "Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG SM-A310F/A310FXXU2BQB1 Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/5.4 Chrome/51.0.2704.106 Mobile Safari/537.36",
			want:      audit.AgentWeb,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := agentTypeFromUserAgent(tt.userAgent)
			assert.Equal(t, tt.want, got, tt.userAgent)
		})
	}
}

func Test_sourceIP(t *testing.T) {
	type args struct {
		XForwardedFor []string
		RemoteAddr    string
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{
			name: "未知",
		},
		{
			name: "直连",
			args: args{
				RemoteAddr: "10.0.2.100:32968",
			},
			want: net.ParseIP("10.0.2.100"),
		},
		{
			name: "经过反向代理",
			args: args{
				XForwardedFor: []string{
					"10.4.37.39, 10.4.71.66",
					"10.110.111.23, 10.110.111.57",
					"142.251.10.100, 142.251.10.138, 142.251.10.102, 142.251.10.101, 142.251.10.113, 142.251.10.139",
				},
			},
			want: net.ParseIP("10.4.37.39"),
		},
		{
			name: "X-Forwarded-For 包含不合法的 IP",
			args: args{
				XForwardedFor: []string{"1.2.3, , 10.4.37.39"},
			},
			want: net.ParseIP("10.4.37.39"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare http request
			req := &http.Request{
				Header:     make(http.Header),
				RemoteAddr: tt.args.RemoteAddr,
			}
			for _, v := range tt.args.XForwardedFor {
				req.Header.Add("X-Forwarded-For", v)
			}

			got := sourceIP(req)
			assert.Equal(t, tt.want, got)
		})
	}
}
