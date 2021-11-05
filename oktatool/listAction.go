package oktatool

import (
	"fmt"
	"github.com/aerostatka/third-party-integrations/logger"
	"github.com/aerostatka/third-party-integrations/models"
)

type ListAction struct {
	repository Repository
	storage    Storage
	parameters *models.OktaToolParameters
	logger     logger.Logger
}

func CreateListAction(rep Repository, st Storage, params *models.OktaToolParameters, log logger.Logger) *ListAction {
	return &ListAction{
		repository: rep,
		storage:    st,
		parameters: params,
		logger:     log,
	}
}

func (action *ListAction) ApplyAction() *models.ActionResult {
	_, err := action.listApplications()
	if err != nil {
		return models.CreateErrorResult(err.Error())
	}

	return models.CreateSuccessfulResult(
		fmt.Sprintf("Application list is stored in %s", action.parameters.DataLocation),
	)
}

func (action *ListAction) listApplications() ([]models.SimpleApp, error) {
	action.logger.Info("Fetching applications....")
	apps, err := action.repository.GetApplications(action.parameters.OnlyActive, action.parameters.Limit)

	if err != nil {
		return apps, err
	}

	action.logger.Info("Done.")
	action.logger.Info(fmt.Sprintf("%d applications were found", len(apps)))

	action.logger.Info("Storing applications to a file")
	err = action.storage.StoreApplicationData(action.parameters.DataLocation, apps)

	if err != nil {
		return apps, err
	}

	action.logger.Info("Done.")

	return apps, nil
}
