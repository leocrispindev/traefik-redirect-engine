package traefik_redirect_engine_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	eng "github.com/leocrispindev/traefik-redirect-engine"
)

func TestDemo(t *testing.T) {
	config := eng.CreateConfig()
	config.IsRedirectEnable = true
	config.Source = "file"
	config.FilePath = "rules.json"

	ctx := context.Background()
	next := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})

	handler, err := eng.New(ctx, next, config, "traefik-redirect-engine")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(recorder, req)

}
