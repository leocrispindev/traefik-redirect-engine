package service

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/leocrispindev/traefik-redirect-engine/internal/model"
	"github.com/leocrispindev/traefik-redirect-engine/internal/utils"
)

func GetDestinyUrl(rule model.Rule, req *http.Request) string {
	fullUrl := utils.GetFullUrl(req)

	result := strings.ReplaceAll(fullUrl, req.Host, rule.Destiny)

	return result
}

func GetRules(config *model.Config) map[string]model.Rule {
	if !config.IsRedirectEnable {
		return make(map[string]model.Rule)
	}
	switch config.Source {
	case "file":
		return GetRulesFromFile(config.FilePath)
	}

	return make(map[string]model.Rule)
}

func GetRulesFromFile(filePath string) map[string]model.Rule {
	result := make(map[string]model.Rule)

	body, err := utils.Readfile(filePath)
	if err != nil {
		os.Stderr.WriteString("Error on read rules file: " + err.Error())
		return result
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		os.Stderr.WriteString("Error on unmarshal rules file " + err.Error())

		return result
	}

	os.Stdout.WriteString("Success loaded rules")

	return result
}

func StartUpdateRedirectRulesJob(config *model.Config, redirectRules map[string]model.Rule) {
	go func() {

		for {
			updateRedirectRules(redirectRules, GetRules(config))
			time.Sleep(30 * time.Second)
		}
	}()

}

func updateRedirectRules(redirectRules map[string]model.Rule, newRules map[string]model.Rule) {
	for key, value := range newRules {
		redirectRules[key] = value
	}
}
