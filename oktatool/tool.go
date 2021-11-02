package oktatool

import (
	"context"
	"github.com/aerostatka/third-party-integrations/config"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type ConsoleTool struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
}

func CreateOktaTool(appConfig *config.AppConfig, action string, actionParameters []string) (*ConsoleTool, error) {
	config := &models.OktaToolParameters{
		Action: action,
	}

	err := config.LoadParameters(actionParameters)

	if err != nil {
		return nil, err
	}

	ctx, client, err := okta.NewClient(
		context.Background(),
		okta.WithOrgUrl(appConfig.Okta.Domain),
		okta.WithToken(appConfig.Okta.Token),
	)

	if err != nil {
		return nil, err
	}

	return &ConsoleTool{
		repository: CreateOktaRepository(ctx, client),
		storage:    CreateLocalFileStorage(),
		parameters: config,
	}, nil
}

func (tool *ConsoleTool) PerformAction() *models.ActionResult {
	switch tool.parameters.Action {
	case models.ActionTypeOktaApplicationList:
		_, err := tool.listApplications()
		if err != nil {
			return models.CreateErrorResult(err.Error())
		}

		return models.CreateSuccessfulResult("Application list is stored in " + *tool.parameters.DataLocation)
	}

	return models.CreateErrorResult("Action is not supported")
}

func (tool *ConsoleTool) listApplications() ([]models.SimpleApp, error) {
	apps, err := tool.repository.GetApplications()

	if err != nil {
		return apps, err
	}

	err = tool.storage.StoreApplicationData(*tool.parameters.DataLocation, apps)

	if err != nil {
		return apps, err
	}

	return apps, nil
}
