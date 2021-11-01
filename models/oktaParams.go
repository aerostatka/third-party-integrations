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
	case ActionTypeOktaApplicationCreate:
		toolParams.DataLocation = &(params[0])
		toolParams.TemplateLocation = &(params[1])
		return nil
	case ActionTypeOktaApplicationList,
		ActionTypeOktaApplicationDisable,
		ActionTypeOktaApplicationEnable,
		ActionTypeOktaApplicationDelete:
		toolParams.DataLocation = &(params[0])
		return nil
	case ActionTypeOktaApplicationUpdateCert:
		toolParams.DataLocation = &(params[0])
		toolParams.TemplateLocation = &(params[1])
		toolParams.CertLocation = &(params[2])
		return nil
	}

	return errors.New("Action " + toolParams.Action + " is not supported")
}
