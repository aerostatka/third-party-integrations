package models

type ActionResult struct {
	Message string
	Error   bool
}

func CreateErrorResult(message string) *ActionResult {
	return &ActionResult{
		Message: message,
		Error:   true,
	}
}

func CreateSuccessfulResult(message string) *ActionResult {
	return &ActionResult{
		Message: message,
		Error:   false,
	}
}

func (result *ActionResult) HasError() bool {
	return result.Error
}
