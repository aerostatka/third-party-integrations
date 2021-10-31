package tools

import (
	"context"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type OktaTool struct {
	context context.Context
	client okta.Client
}

/*func (tool *OktaTool) PerformAction() (ActionResult, error) {

}

func CreateOktaTool() (*OktaTool, error) {

}*/