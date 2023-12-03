package service

import (
	"net/http"
	"strings"
	"traefik-redirect-engine/internal/model"
	"traefik-redirect-engine/internal/utils"
)

func GetDestinyUrl(rule model.Rule, req *http.Request) string {
	fullUrl := utils.GetFullUrl(req)

	result := strings.ReplaceAll(fullUrl, req.Host, rule.Destiny)

	return result
}
