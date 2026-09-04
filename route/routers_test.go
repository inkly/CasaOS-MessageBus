package route

import (
	"net/http"
	"testing"
)

func TestSkipAccessLog(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		realIP   string
		host     string
		expected bool
	}{
		{"utilization publish from loopback IPv4", http.MethodPost, "/v2/message_bus/event/casaos/casaos:system:utilization", "127.0.0.1", "127.0.0.1:43251", true},
		{"utilization publish from loopback IPv6", http.MethodPost, "/v2/message_bus/event/casaos/casaos:system:utilization", "::1", "[::1]:43251", true},
		{"publish over the unix socket", http.MethodPost, "/v2/message_bus/event/app-management/app:install:progress", "", "unix", true},
		{"publish from the LAN stays logged", http.MethodPost, "/v2/message_bus/event/casaos/casaos:system:utilization", "192.168.1.20", "nas.local", false},
		{"websocket subscribe stays logged", http.MethodGet, "/v2/message_bus/event/casaos", "127.0.0.1", "127.0.0.1:1", false},
		{"event type registration stays logged", http.MethodPost, "/v2/message_bus/event_type", "127.0.0.1", "127.0.0.1:1", false},
		{"action publish stays logged", http.MethodPost, "/v2/message_bus/action/casaos/x", "127.0.0.1", "127.0.0.1:1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := skipAccessLog(tt.method, tt.path, tt.realIP, tt.host); actual != tt.expected {
				t.Errorf("skipAccessLog(%q, %q, %q, %q) = %v, want %v", tt.method, tt.path, tt.realIP, tt.host, actual, tt.expected)
			}
		})
	}
}
