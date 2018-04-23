package Commands

import "SectionInformations"
import "errors"
import "strings"

// appName TEXT
type AppNameCommand struct {
	commandId uint8
	name string
}

func NewAppNameCommand() *AppNameCommand {
	return &AppNameCommand{1, "appName"}
}

// [command id:1 byte][parameter:x bytes]
func (cmd *AppNameCommand) GetByteCode(parameters []string, sectionName string, labelName string, sinfo *SectionInformations.SectionInformations) ([]byte, error) {
	retByteCode := make([]byte, 0)
	if !cmd.isParametersCountCorrect(parameters) {
		if len(labelName) > 0 {
			return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName + ", label: " + labelName)
		}
		return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName)
	}

	convertedTextParameter := []byte(parameters[0])
	convertedTextParameter = append(convertedTextParameter, byte(0)) // Make it asciiz

	retByteCode = append(retByteCode, cmd.commandId)
	retByteCode = append(retByteCode, convertedTextParameter...)
	return retByteCode, nil
}

func (cmd *AppNameCommand) isParametersCountCorrect(parameters []string) bool {
	return len(parameters) == 1
}