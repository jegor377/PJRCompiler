package Commands

import "pjre/SectionInformations"
import "errors"
import "strings"

// wydt TEXT
type WYDTCommand struct {
	commandId uint8
	name string
}

func NewWYDTCommand() *WYDTCommand {
	return &WYDTCommand{2, "wydt"}
}

// [command id:1 byte][parameter:x bytes]
func (cmd *WYDTCommand) GetByteCode(parameters []string, sectionName string, labelName string, sinfo *SectionInformations.SectionInformations) ([]byte, error) {
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

func (cmd *WYDTCommand) isParametersCountCorrect(parameters []string) bool {
	return len(parameters) == 1
}