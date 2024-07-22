package plugin_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	plugin "github.com/RiskIdent/traefik-tls-headers-plugin"
)

func TestInvalidConfig(t *testing.T) {
	cfg := plugin.CreateConfig()
	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	_, err := plugin.New(context.Background(), next, cfg, "traefik-tls-headers-plugin")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTLSCipher(t *testing.T) {
	cfg := plugin.CreateConfig()
	cfg.Headers.Cipher = "X-TLS-Cipher"
	next := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		assertHeader(t, r.Header, "X-TLS-Cipher", "TLS_AES_128_GCM_SHA256")
	})
	handler, err := plugin.New(context.Background(), next, cfg, "traefik-tls-headers-plugin")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewTLSServer(handler)
	defer server.Close()

	req, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	client := server.Client()
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(resp.Body)
}

func assertHeader(t *testing.T, header http.Header, key, expected string) {
	t.Helper()

	if header.Get(key) != expected {
		t.Errorf("invalid header value\nwant: %s=%q\ngot:  %s=%q", key, expected, key, header.Get(key))
	}
}
