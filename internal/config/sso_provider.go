package config

import "github.com/pkg/errors"

const (
	Google = "google"
)

type SSOProvider interface {
	Credentials(string) ssoProviderConfig
}

type ssoProvider struct {
	providerConfig yamlSSOProviderConfig
}

type ssoProviderConfig struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	RedirectURI  string `yaml:"redirect_uri"`
}

type yamlSSOProviderConfig struct {
	Google ssoProviderConfig `yaml:"google"`
}

func NewSSOProvider(providersCreds yamlSSOProviderConfig) SSOProvider {
	return &ssoProvider{
		providerConfig: providersCreds,
	}
}

func (s ssoProvider) Credentials(providerName string) ssoProviderConfig {
	switch providerName {
	case Google:
		return s.providerConfig.Google
	default:
		panic(errors.Errorf("unknown provider %s", providerName))
	}
}
