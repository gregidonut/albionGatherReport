package gatherReport

import "errors"

var (
	parsingTemplateErr = errors.New("error parsing template")
	creatingReadMeErr  = errors.New("Error creating README.md:")
	writingReadMeErr   = errors.New("Error writing to README.md:")
	toJSONErr          = errors.New("error converting to json")
	marshalErr         = errors.New("error marshalling struct")
	parsingKwargsErr   = errors.New("error parsing kwargs")
)
