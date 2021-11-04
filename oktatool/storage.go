package oktatool

import (
	"encoding/csv"
	"errors"
	"github.com/aerostatka/third-party-integrations/models"
	"os"
	"path"
)

var (
	localCsvFileStorageSupportedExtensions = map[string]bool{
		".csv": true,
		".CSV": true,
	}
)

type Storage interface {
	GetAppsData(location string) ([]models.SimpleApp, error)
	GetTemplate(location string, templateName string) error
	StoreApplicationData(location string, apps []models.SimpleApp) error
}

type LocalCsvFileStorage struct {
}

func CreateLocalCsvFileStorage() *LocalCsvFileStorage {
	return &LocalCsvFileStorage{}
}

func (storage *LocalCsvFileStorage) validateLocation(location string) bool {
	ext := path.Ext(location)
	return localCsvFileStorageSupportedExtensions[ext]
}

func (storage *LocalCsvFileStorage) GetAppsData(location string) ([]models.SimpleApp, error) {
	var apps []models.SimpleApp
	return apps, nil
}

func (storage *LocalCsvFileStorage) GetTemplate(location string, templateName string) error {
	return nil
}

func (storage *LocalCsvFileStorage) StoreApplicationData(location string, apps []models.SimpleApp) error {
	if len(apps) == 0 {
		return nil
	}

	if !storage.validateLocation(location) {
		return errors.New("Extension is not supported")
	}

	csvFile, err := os.Create(location)
	if err != nil {
		return err
	}

	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	line := []string{"Id", "Code", "Label", "Status"}
	err = writer.Write(line)

	if err != nil {
		writer.Flush()
		return err
	}

	for _, app := range apps {
		line := []string{app.Id, app.Code, app.Label, app.Status}

		err = writer.Write(line)

		if err != nil {
			writer.Flush()
			return err
		}
	}

	writer.Flush()

	return nil
}
