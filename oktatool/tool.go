package oktatool

import (
	"context"
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

type ConsoleAction interface {
	ApplyAction() *models.ActionResult
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
		storage:    CreateLocalCsvFileStorage(),
		parameters: config,
		logger:     log,
	}, nil
}

func (tool *ConsoleTool) PerformAction() *models.ActionResult {
	switch tool.parameters.Action {
	case models.ActionTypeOktaApplicationList:
		action := CreateListAction(tool.repository, tool.storage, tool.parameters, tool.logger)

		return action.ApplyAction()
	case models.ActionTypeOktaApplicationDisable:
		action := CreateDisableAction(tool.repository, tool.storage, tool.parameters, tool.logger)

		return action.ApplyAction()
	case models.ActionTypeOktaApplicationEnable:
		action := CreateEnableAction(tool.repository, tool.storage, tool.parameters, tool.logger)

		return action.ApplyAction()
	}

	return models.CreateErrorResult("Action is not supported")
}
