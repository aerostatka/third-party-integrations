package models

import (
	"errors"
	"net/url"
)

type SimpleApp struct {
	Id     string
	Code   string
	Url    string
	Label  string
	Status string
}

func (app SimpleApp) FindStorageAppInList(list []SimpleApp) *SimpleApp {
	for _, listApp := range list {
		if app.Label == listApp.Label {
			listApp.Url = app.Url
			return &listApp
		}
	}

	return nil
}

func (app SimpleApp) ValidateForCreation() error {
	if app.Url == "" {
		return errors.New("URL cannot be empty")
	}

	u, err := url.Parse(app.Url)
	if err != nil {
		return err
	}

	if u.Host == "" {
		return errors.New("URL should have a schema")
	}

	return nil
}
