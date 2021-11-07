package models

import (
	"errors"
	"fmt"
)

const (
	ToolTypeOkta string = "okta"
)

const (
	ActionTypeOktaApplicationList       string = "applicationList"
	ActionTypeOktaApplicationCreate     string = "applicationCreate"
	ActionTypeOktaApplicationDisable    string = "applicationDisable"
	ActionTypeOktaApplicationEnable     string = "applicationEnable"
	ActionTypeOktaApplicationDelete     string = "applicationDelete"
	ActionTypeOktaApplicationUpdateCert string = "applicationUpdateCert"
)

type FactoryParams struct {
	ToolType         string
	ActionType       string
	ActionParameters []string
}

func ConvertToFactoryParams(params []string) (*FactoryParams, error) {
	if len(params) < 2 {
		return nil, errors.New("Parameters count is less than 2")
	}

	factoryParams := &FactoryParams{
		ToolType:         params[0],
		ActionType:       params[1],
		ActionParameters: params[2:],
	}

	err := factoryParams.validate()

	if err != nil {
		return nil, err
	}

	return factoryParams, nil
}

func (factoryParams *FactoryParams) hasValidToolType() bool {
	return factoryParams.ToolType == ToolTypeOkta
}

func (factoryParams *FactoryParams) hasValidActionType() bool {
	availableOktaActions := map[string]bool{
		ActionTypeOktaApplicationList:       true,
		ActionTypeOktaApplicationCreate:     true,
		ActionTypeOktaApplicationDisable:    true,
		ActionTypeOktaApplicationEnable:     true,
		ActionTypeOktaApplicationDelete:     true,
		ActionTypeOktaApplicationUpdateCert: true,
	}

	switch factoryParams.ToolType {
	case ToolTypeOkta:
		_, ok := availableOktaActions[factoryParams.ActionType]
		return ok
	}

	return false
}

func (factoryParams *FactoryParams) hasValidParamsNumber() bool {
	switch factoryParams.ToolType {
	case ToolTypeOkta:
		switch factoryParams.ActionType {
		case ActionTypeOktaApplicationCreate:
			return len(factoryParams.ActionParameters) > 1 && len(factoryParams.ActionParameters) < 5
		case ActionTypeOktaApplicationList:
			return len(factoryParams.ActionParameters) > 0 && len(factoryParams.ActionParameters) < 4
		case ActionTypeOktaApplicationDisable,
			ActionTypeOktaApplicationEnable,
			ActionTypeOktaApplicationDelete:
			return len(factoryParams.ActionParameters) == 1
		case ActionTypeOktaApplicationUpdateCert:
			return len(factoryParams.ActionParameters) == 2
		}
	}

	return false
}

func (factoryParams *FactoryParams) validate() error {
	if !factoryParams.hasValidToolType() {
		return errors.New(fmt.Sprintf("Tool type %s is not supported", factoryParams.ToolType))
	}

	if !factoryParams.hasValidActionType() {
		return errors.New(fmt.Sprintf(
			"Action type %s is not supported for tool %s",
			factoryParams.ActionType,
			factoryParams.ToolType),
		)
	}

	if !factoryParams.hasValidParamsNumber() {
		return errors.New(fmt.Sprintf(
			"Action type %s for tool %s has the wrong number of parameters",
			factoryParams.ActionType,
			factoryParams.ToolType),
		)
	}

	return nil
}
