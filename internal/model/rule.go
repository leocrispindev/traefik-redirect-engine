package model

type Rule struct {
	RedirectUrl string `json:"url"`
	URIs        []UriRule
}

type UriRule struct {
	UriPath        string `json:"uriPath"`
	URLRedirectURI string `json:"url"`
}
