package plugin

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidConfig(t *testing.T) {
	cfg := CreateConfig()
	next := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	_, err := New(context.Background(), next, cfg, "traefik-tls-headers-plugin")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestTLSVersion(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers.Version = "X-Tls-Version"
	next := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		assertHeader(t, r.Header, "X-Tls-Version", "TLS 1.3")
	})
	handler, err := New(context.Background(), next, cfg, "traefik-tls-headers-plugin")
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

func TestTLSCipher(t *testing.T) {
	cfg := CreateConfig()
	cfg.Headers.Cipher = "X-Tls-Cipher"
	next := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		assertHeader(t, r.Header, "X-Tls-Cipher", "TLS_AES_128_GCM_SHA256")
	})
	handler, err := New(context.Background(), next, cfg, "traefik-tls-headers-plugin")
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
