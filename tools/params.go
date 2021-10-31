package tools

import (
	"errors"
)

const (
	ToolTypeOkta string = "okta"
)

const (
	ActionTypeOktaApplicationList string = "applicationList"
	ActionTypeOktaApplicationCreate string = "applicationCreate"
	ActionTypeOktaApplicationDisable string = "applicationDisable"
	ActionTypeOktaApplicationEnable string = "applicationEnable"
	ActionTypeOktaApplicationDelete string = "applicationDelete"
)

type FactoryParams struct {
	ToolType string
	ActionType string
	ActionParameters []string
}

func (factoryParams *FactoryParams) hasValidToolType() bool {
	return factoryParams.ToolType == ToolTypeOkta
}

func (factoryParams *FactoryParams) hasValidActionType() bool {
	availableOktaActions := map[string]bool {
		ActionTypeOktaApplicationList: true,
		ActionTypeOktaApplicationCreate: true,
		ActionTypeOktaApplicationDisable: true,
		ActionTypeOktaApplicationEnable: true,
		ActionTypeOktaApplicationDelete: true,
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
				case ActionTypeOktaApplicationList:
					return len(factoryParams.ActionParameters) == 0
				case ActionTypeOktaApplicationCreate,
				ActionTypeOktaApplicationDisable,
				ActionTypeOktaApplicationEnable,
				ActionTypeOktaApplicationDelete:
					return len(factoryParams.ActionParameters) == 1
			}
	}

	return false
}

func (factoryParams *FactoryParams) Validate() error {
	if !factoryParams.hasValidToolType() {
		return errors.New("Tool type " + factoryParams.ToolType + " is not supported")
	}

	if !factoryParams.hasValidActionType() {
		return errors.New("Action type " + factoryParams.ActionType +
			" is not supported for tool " + factoryParams.ToolType)
	}

	if !factoryParams.hasValidParamsNumber() {
		return errors.New("Action type " + factoryParams.ActionType +
			" for tool " + factoryParams.ToolType +
			" has the wrong number of parameters")
	}

	return nil
}

func ConvertToFactoryParams(params []string) (*FactoryParams, error)  {
	if len(params) < 2 {
		return nil, errors.New("Parameters count is less than 2")
	}

	return &FactoryParams{
		ToolType: params[0],
		ActionType: params[1],
		ActionParameters: params[2:],
	}, nil
}