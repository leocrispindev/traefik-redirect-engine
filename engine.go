package redirectengine

import (
	"context"
	"net/http"
	model "traefik-redirect-engine/internal/model"
)

type Config struct {
	IsRedirectEnable bool   `json:"isRedirectEnable"`
	Source           string `json:"source"`
}

// plugin configuration
func CreateConfig() *Config {
	return &Config{}
}

type Engine struct {
	next          http.Handler
	isEnable      bool
	Regex         string
	RedirectRules map[string]model.Rule
}

// create plugin instance
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &Engine{
		next:     next,
		isEnable: config.IsRedirectEnable,
	}, nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	if !e.isEnable {
		e.next.ServeHTTP(rw, req)
	}

	e.next.ServeHTTP(rw, req)
}
