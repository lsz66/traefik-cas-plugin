// Package traefik_cas_plugin a cas plugin.
package traefik_cas_plugin

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	Url string
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Url: "",
	}
}

// CAS a CAS plugin.
type CAS struct {
	next http.Handler
	name string
	url  string
}

// New created a new CAS plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Url) == 0 {
		return nil, fmt.Errorf("CAS Server URL cannot be empty")
	}

	return &CAS{
		next: next,
		name: name,
		url:  config.Url,
	}, nil
}

func (r *CAS) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for _, cookie := range req.Cookies() {
		if strings.EqualFold(cookie.Name, "SESSION") {
			r.next.ServeHTTP(rw, req)
			break
		}
	}
	http.Redirect(rw, req, r.url, 301)
}
