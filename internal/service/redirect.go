package service

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	log "github.com/leocrispindev/traefik-redirect-engine/internal/log"
	"github.com/leocrispindev/traefik-redirect-engine/internal/model"
	"github.com/leocrispindev/traefik-redirect-engine/internal/utils"
)

func GetDestinyUrl(rule model.Rule, req *http.Request) string {
	fullUrl := utils.GetFullUrl(req)

	result := strings.Replace(fullUrl, req.Host, rule.Destiny, 1)

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
		log.Error("Error on read rules file: " + err.Error())
		return result
	}

	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Error("Error on unmarshal rules file " + err.Error())

		return result
	}

	log.Info("Success loaded rules from file")

	return result
}

func StartUpdateRedirectRulesJob(config *model.Config, redirectRules map[string]model.Rule) {
	go func() {
		log.Info("Started Update rules job")

		for {
			updateRedirectRules(redirectRules, GetRules(config))
			time.Sleep(30 * time.Second)
		}
	}()

}

func updateRedirectRules(redirectRules map[string]model.Rule, newRules map[string]model.Rule) {
	log.Info("Start update rules")
	for key, value := range newRules {
		redirectRules[key] = value
	}

	log.Info("Finish update rules")

}
