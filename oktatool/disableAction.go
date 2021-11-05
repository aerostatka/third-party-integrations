package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type DisableAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateDisableAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *DisableAction {
	return &DisableAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *DisableAction) ApplyAction() *models.ActionResult {
	err := action.disableApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were disabled", action.parameters.DataLocation),
	)
}

func (action *DisableAction) disableApplications() error {
	action.logger.Info("Fetching applications from storage....")
	apps, err := action.storage.GetAppsData(action.parameters.DataLocation)
	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications successfully fetched", len(apps)))
	action.logger.Info("Fetching applications from OKTA....")

	existingApps, err := action.repository.GetApplications(true, action.parameters.Limit)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%d active applications were found", len(existingApps)))

	for _, app := range apps {
		action.logger.Info(fmt.Sprintf("Processing %s application.", app.Label))
		oktaApp := app.FindStorageAppInList(existingApps)

		if oktaApp == nil {
			action.logger.Info(fmt.Sprintf("Application %s is not found in OKTA.", app.Label))
		} else {
			action.logger.Info(fmt.Sprintf("Application %s is found in OKTA.", app.Label))
			// Implement application disable
		}
	}

	return nil
}
