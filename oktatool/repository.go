package oktatool

import (
	"context"
	"github.com/aerostatka/third-party-integrations/models"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type Repository interface {
	GetApplications() ([]models.SimpleApp, error)
}

type OktaRepository struct {
	context context.Context
	client  *okta.Client
}

func CreateOktaRepository(ctx context.Context, ct *okta.Client) *OktaRepository {
	return &OktaRepository{
		context: ctx,
		client:  ct,
	}
}

func (rep *OktaRepository) GetApplications() ([]models.SimpleApp, error) {
	var apps []models.SimpleApp
	_, _, err := rep.client.Application.ListApplications(rep.context, nil)

	if err != nil {
		return apps, err
	}

	return apps, nil
}
