package traefik_redirect_engine

import (
	"context"
	"net/http"

	log "github.com/leocrispindev/traefik-redirect-engine/internal/log"
	model "github.com/leocrispindev/traefik-redirect-engine/internal/model"
	"github.com/leocrispindev/traefik-redirect-engine/internal/service"
)

// plugin configuration
func CreateConfig() *model.Config {
	return &model.Config{}
}

type Engine struct {
	next          http.Handler
	name          string
	isEnable      bool
	Regex         string
	RedirectRules map[string]model.Rule
	FilePath      string
}

// create plugin instance
func New(ctx context.Context, next http.Handler, config *model.Config, name string) (http.Handler, error) {
	// ...
	result := Engine{
		next:          next,
		name:          "traefik redirect engine",
		isEnable:      config.IsRedirectEnable,
		FilePath:      config.FilePath,
		RedirectRules: service.GetRules(config),
	}

	go service.StartUpdateRedirectRulesJob(config, result.RedirectRules)

	return &result, nil
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// ...
	if !e.isEnable {
		e.next.ServeHTTP(rw, req)
	}

	rule, found := e.RedirectRules[req.Host]
	if !found {
		e.next.ServeHTTP(rw, req)
	}

	destinyURL := service.GetDestinyUrl(rule, req)

	log.Info("Rule found, redirect destiny: " + destinyURL)

	http.Redirect(rw, req, destinyURL, http.StatusPermanentRedirect)
	//e.next.ServeHTTP(rw, req)
}
