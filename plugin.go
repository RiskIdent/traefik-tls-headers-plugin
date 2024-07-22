// Package plugin contains the Traefik plugin for adding headers based on the
// TLS information
package plugin

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
)

var errMissingHeaderConfig = errors.New("missing header config: must set headers.cipher")

// Config the plugin configuration.
type Config struct {
	Headers ConfigHeaders `json:"headers,omitempty"`
}

// ConfigHeaders defines the headers to use for the different values.
type ConfigHeaders struct {
	Cipher string `json:"port,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: ConfigHeaders{},
	}
}

// TLSHeadersPlugin is the main handler model for this Traefik plugin.
type TLSHeadersPlugin struct {
	next    http.Handler
	headers ConfigHeaders
	name    string
}

// New created a new TLSHeadersPlugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Headers == (ConfigHeaders{}) {
		return nil, errMissingHeaderConfig
	}

	return &TLSHeadersPlugin{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *TLSHeadersPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if a.headers.Cipher != "" && req.TLS != nil {
		req.Header.Set(a.headers.Cipher, tls.CipherSuiteName(req.TLS.CipherSuite))
	}

	a.next.ServeHTTP(rw, req)
}
