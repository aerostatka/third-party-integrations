package oktatool

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
)

var (
	localFileStorageSupportedExtensions = map[string]bool{
		".csv": true,
		".CSV": true,
	}

	localFileStorageSupportedJsonExtensions = map[string]bool{
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
	GetTemplate(location string, templateName string) (*models.Template, error)
	StoreApplicationsData(location string, apps []models.SimpleApp) error
	StoreApplicationData(location string, app okta.App) error
}
type LocalCsvFileStorage struct {
}

func CreateLocalCsvFileStorage() *LocalCsvFileStorage {
	return &LocalCsvFileStorage{}
}

func (storage *LocalCsvFileStorage) validateFileLocation(location string) bool {
	ext := path.Ext(location)
	return localFileStorageSupportedExtensions[ext]
}

func (storage *LocalCsvFileStorage) validateJsonFileLocation(location string) bool {
	ext := path.Ext(location)
	return localFileStorageSupportedJsonExtensions[ext]
}

func (storage *LocalCsvFileStorage) GetAppsData(location string) ([]models.SimpleApp, error) {
	var apps []models.SimpleApp

	if !storage.validateFileLocation(location) {
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

func (storage *LocalCsvFileStorage) StoreApplicationsData(location string, apps []models.SimpleApp) error {
	if len(apps) == 0 {
		return nil
	}

	if !storage.validateFileLocation(location) {
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

func (storage *LocalCsvFileStorage) StoreApplicationData(location string, app okta.App) error {
	if !storage.validateJsonFileLocation(location) {
		return errors.New("Extension is not supported")
	}

	jsonString, err := json.MarshalIndent(app, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(location, jsonString, fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func (storage *LocalCsvFileStorage) GetTemplate(location string, templateName string) (*models.Template, error) {
	if !storage.validateJsonFileLocation(location) {
		return nil, errors.New("Extension is not supported")
	}

	jsonFile, err := os.Open(location)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var templates models.Templates
	err = json.Unmarshal(byteValue, &templates)
	if err != nil {
		return nil, err
	}

	template, ok := templates.Templates[templateName]
	if !ok {
		return nil, errors.New("Template is not found")
	}

	return &template, nil
}
