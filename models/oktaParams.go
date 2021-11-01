package models

import "errors"

type OktaToolParameters struct {
	Action           string
	DataLocation     *string
	TemplateLocation *string
	CertLocation     *string
}

func (toolParams *OktaToolParameters) LoadParameters(params []string) error {
	switch toolParams.Action {
	case ActionTypeOktaApplicationList:
		return nil
	case ActionTypeOktaApplicationCreate,
		ActionTypeOktaApplicationDisable,
		ActionTypeOktaApplicationEnable,
		ActionTypeOktaApplicationDelete:
		toolParams.DataLocation = &(params[0])
		return nil
	}

	return errors.New("Action " + toolParams.Action + " is not supported")
}
