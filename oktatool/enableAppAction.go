package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type EnableAppAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateEnableAppAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *EnableAppAction {
	return &EnableAppAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *EnableAppAction) ApplyAction() *models.ActionResult {
	err := action.enableApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Found applications from the list %s were enabled", action.parameters.DataLocation),
	)
}

func (action *EnableAppAction) enableApplications() error {
	action.logger.Info("Fetching applications from storage....")
	apps, err := action.storage.GetAppsData(action.parameters.DataLocation)
	if err != nil {
		return err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications successfully fetched", len(apps)))
	action.logger.Info("Fetching applications from OKTA....")

	inactiveApps, err := action.repository.GetApplications(
		models.OktaParamsApplicationsStatusInactive,
		action.parameters.Limit,
	)
	if err != nil {
		return err
	}

	action.logger.Info(fmt.Sprintf("%d inactive applications were found", len(inactiveApps)))

	for _, app := range apps {
		action.logger.Info(fmt.Sprintf("Processing %s application.", app.Label))
		oktaApp := app.FindStorageAppInList(inactiveApps)

		if oktaApp == nil {
			action.logger.Info(fmt.Sprintf("Inactive application %s is not found in OKTA.", oktaApp.Label))
		} else {
			action.logger.Info(fmt.Sprintf("Inactive application %s is found in OKTA.", oktaApp.Label))
			action.logger.Info("Enabling application....")
			err := action.repository.ChangeApplicationStatus(oktaApp, models.OktaParamsApplicationsStatusActive)
			if err != nil {
				return err
			}
			action.logger.Info("Done.")
		}
	}

	return nil
}
