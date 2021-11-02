package oktatool

import "github.com/aerostatka/third-party-integrations/models"

type Storage interface {
	GetAppsData(location string) ([]models.SimpleApp, error)
	GetTemplate(location string, templateName string) error
	StoreApplicationData(location string, apps []models.SimpleApp) error
}

type LocalFileStorage struct {
}

func CreateLocalFileStorage() *LocalFileStorage {
	return &LocalFileStorage{}
}

func (storage *LocalFileStorage) GetAppsData(location string) ([]models.SimpleApp, error) {
	var apps []models.SimpleApp
	return apps, nil
}

func (storage *LocalFileStorage) GetTemplate(location string, templateName string) error {
	return nil
}

func (storage *LocalFileStorage) StoreApplicationData(location string, apps []models.SimpleApp) error {
	return nil
}
