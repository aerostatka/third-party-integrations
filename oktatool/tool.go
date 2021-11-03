package oktatool

import (
	"context"
	"fmt"
	"github.com/aerostatka/third-party-integrations/config"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type ConsoleTool struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateOktaTool(appConfig *config.AppConfig, action string, actionParameters []string, log logger.Logger) (*ConsoleTool, error) {
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
		okta.WithRequestTimeout(30),
		okta.WithRateLimitMaxRetries(1),
		okta.WithConnectionTimeout(15),
	)

	if err != nil {
		return nil, err
	}

	return &ConsoleTool{
		repository: CreateOktaRepository(ctx, client),
		storage:    CreateLocalFileStorage(),
		parameters: config,
		logger:     log,
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

	tool.logger.Info(fmt.Sprintf("%d applications were found", len(apps)))

	err = tool.storage.StoreApplicationData(*tool.parameters.DataLocation, apps)

	if err != nil {
		return apps, err
	}

	return apps, nil
}
