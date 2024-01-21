package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	log "github.com/leocrispindev/traefik-redirect-engine/internal/log"
	"github.com/leocrispindev/traefik-redirect-engine/internal/model"
	"github.com/leocrispindev/traefik-redirect-engine/internal/utils"
)

var RedirectUriRegexMap = make(map[string][]*regexp.Regexp)

func GetDestinyUrl(rule model.Rule, req *http.Request) string {
	fullUrl := utils.GetFullUrl(req)

	redirectUrlFromURI, found := hasUriRule(rule, req.RequestURI)

	if found {
		return replaceUrl(fullUrl, req.Host, redirectUrlFromURI)
	}

	return replaceUrl(fullUrl, req.Host, rule.RedirectUrl)

}

func replaceUrl(fullUrl string, host string, redirectUrl string) string {
	return strings.Replace(fullUrl, host, redirectUrl, 1)

}

func hasUriRule(rule model.Rule, uri string) (string, bool) {
	uriRulesByHost := RedirectUriRegexMap[rule.RedirectUrl]

	if uriRulesByHost == nil {
		return "", false
	}

	for i := 0; i < len(uriRulesByHost)-1; i++ {
		match, err := matchUri(uriRulesByHost[i], uri)

		if err != nil || !match {
			continue
		}

		return rule.URIs[i].URLRedirectURI, true
	}

	return "", false
}

func matchUri(regex *regexp.Regexp, requestUri string) (bool, error) {

	return regex.Match([]byte(requestUri)), nil

}

func CompileAllUriRules(rules map[string]model.Rule) {

	for host, rule := range rules {

		var regex []*regexp.Regexp

		for _, uri := range rule.URIs {
			rgx, err := compileRegex(uri.URLRedirectURI)

			if err != nil {
				continue
			}

			regex = append(regex, rgx)
		}

		RedirectUriRegexMap[host] = regex
	}
}

func compileRegex(uriRule string) (*regexp.Regexp, error) {
	regexString := fmt.Sprintf(`^%s`, regexp.QuoteMeta(uriRule))

	// Compilar a regex
	return regexp.Compile(regexString)
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
	log.Info("Started Update rules job")

	for {
		updateRedirectRules(redirectRules, GetRules(config))
		time.Sleep(30 * time.Second)
	}

}

func updateRedirectRules(redirectRules map[string]model.Rule, newRules map[string]model.Rule) {
	log.Info("Start update rules")
	for key, value := range newRules {
		redirectRules[key] = value
	}

	log.Info("Finish update rules")

}
