package oktatool

import (
	"encoding/csv"
	"errors"
	"github.com/aerostatka/third-party-integrations/models"
	"io"
	"os"
	"path"
)

var (
	localCsvFileStorageSupportedExtensions = map[string]bool{
		".csv": true,
		".CSV": true,
	}

	localCsvFileStorageSupportedTemplateExtensions = map[string]bool{
		".json": true,
		".JSON": true,
		".txt":  true,
		".TXT":  true,
	}
)

const (
	storageApplicationCommonHeader = "Name"
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

func (storage *LocalCsvFileStorage) validateApplicationFileLocation(location string) bool {
	ext := path.Ext(location)
	return localCsvFileStorageSupportedExtensions[ext]
}

func (storage *LocalCsvFileStorage) validateTemplateFileLocation(location string) bool {
	ext := path.Ext(location)
	return localCsvFileStorageSupportedTemplateExtensions[ext]
}

func (storage *LocalCsvFileStorage) GetAppsData(location string) ([]models.SimpleApp, error) {
	var apps []models.SimpleApp

	if !storage.validateApplicationFileLocation(location) {
		return apps, errors.New("Extension is not supported")
	}

	csvFile, err := os.Open(location)
	if err != nil {
		return apps, err
	}

	defer csvFile.Close()
	reader := csv.NewReader(csvFile)

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return apps, err
		}

		length := len(record)

		if length > 0 && record[0] != storageApplicationCommonHeader {
			app := models.SimpleApp{
				Label: record[0],
			}

			if length > 1 {
				app.Url = record[1]
			}

			apps = append(apps, app)
		}
	}

	return apps, nil
}

func (storage *LocalCsvFileStorage) GetTemplate(location string, templateName string) error {
	return nil
}

func (storage *LocalCsvFileStorage) StoreApplicationData(location string, apps []models.SimpleApp) error {
	if len(apps) == 0 {
		return nil
	}

	if !storage.validateApplicationFileLocation(location) {
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
