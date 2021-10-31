package config

import (
	"errors"
	"net/url"
	"os"
)

type OktaConfig struct {
	Domain string
	Token string
}

type AppConfig struct {
	Okta OktaConfig
}

type Config interface {
	GetConfig() *AppConfig
}

type ContextConfig struct {
	config *AppConfig
}

func contextConfigSanityCheck() error {
	if os.Getenv("OKTA_ORG_URL") == "" ||
		os.Getenv("OKTA_TOKEN") == "" {
		return errors.New("OKTA environment variables are not set properly")
	}

	return nil
}

func CreateContextConfig() (*ContextConfig, error) {
	err := contextConfigSanityCheck()

	if err != nil {
		return nil, err
	}

	appConfig := &AppConfig{
		Okta: OktaConfig{
			Domain: os.Getenv("OKTA_ORG_URL"),
			Token: os.Getenv("OKTA_TOKEN"),
		},
	}

	err = appConfig.validate()

	if err != nil {
		return nil, err
	}

	return &ContextConfig{
		config: appConfig,
	}, nil
}

func (conf *ContextConfig) GetConfig() *AppConfig {
	return conf.config
}

func (appConfig *AppConfig) validate() error {
	_, err := url.ParseRequestURI(appConfig.Okta.Domain)

	return err
}
