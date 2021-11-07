package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type DisableAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateDisableAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *DisableAppAction {
	return &DisableAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *DisableAppAction) ApplyAction() *models.ActionResult {
	err := action.disableApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were disabled", action.parameters.DataLocation),
	)
}

func (action *DisableAppAction) disableApplications() error {
	action.logger.Info("Fetching applications from storage....")
	apps, err := action.storage.GetAppsData(action.parameters.DataLocation)
	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications successfully fetched", len(apps)))
	action.logger.Info("Fetching applications from OKTA....")

	activeApps, err := action.repository.GetApplications(
		models.OktaParamsApplicationsStatusActive,
		action.parameters.Limit,
	)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%d active applications were found", len(activeApps)))

	for _, app := range apps {
		action.logger.Info(fmt.Sprintf("Processing %s application.", app.Label))
		oktaApp := app.FindStorageAppInList(activeApps)

		if oktaApp == nil {
			action.logger.Info(fmt.Sprintf("Active application %s is not found in OKTA.", oktaApp.Label))
		} else {
			action.logger.Info(fmt.Sprintf("Active application %s is found in OKTA.", oktaApp.Label))
			action.logger.Info("Disabling application....")
			err := action.repository.ChangeApplicationStatus(oktaApp, models.OktaParamsApplicationsStatusInactive)
			if err != nil {
				return err
			}
			action.logger.Info("Done.")
		}
	}

	return nil
}
