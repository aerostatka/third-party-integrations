package oktatool

import (
	"context"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type ConsoleTool struct {
	Context    context.Context
	Client     *okta.Client
	Parameters *models.OktaToolParameters
}

func (tool *ConsoleTool) PerformAction() (*models.ActionResult, error) {
	return nil, nil
}
