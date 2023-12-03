package redirectengine

import (
	"context"
	"net/http"
	model "traefik-redirect-engine/internal/model"
	"traefik-redirect-engine/internal/service"
)

type Config struct {
	IsRedirectEnable bool   `json:"isRedirectEnable"`
	Source           string `json:"source"`
	FilePath         string `json:"filePath"`
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
	FilePath      string
}

// create plugin instance
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// ...
	return &Engine{
		next:     next,
		isEnable: config.IsRedirectEnable,
		FilePath: config.FilePath,
	}, nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	if !e.isEnable {
		e.next.ServeHTTP(rw, req)
	}

	rule, exists := e.RedirectRules[req.Host]
	if !exists {
		e.next.ServeHTTP(rw, req)
	}

	destinyURL := service.GetDestinyUrl(rule, req)

	http.Redirect(rw, req, destinyURL, http.StatusPermanentRedirect)

	//e.next.ServeHTTP(rw, req)
}
