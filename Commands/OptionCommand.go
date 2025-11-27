package Commands

import "pjre/SectionInformations"
import "errors"
import "strings"
import "pjre/ConvertingTools"

// option dialogOptionName
type OptionCommand struct {
	commandId uint8
	name string
}

func NewOptionCommand() *OptionCommand {
	return &OptionCommand{5 , "option"}
}

// [command id:1 byte][parameter:x bytes]
func (cmd *OptionCommand) GetByteCode(parameters []string, sectionName string, labelName string, sinfo *SectionInformations.SectionInformations) ([]byte, error) {
	retByteCode := make([]byte, 0)
	if !cmd.isParametersCountCorrect(parameters) {
		if len(labelName) > 0 {
			return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName + ", label: " + labelName)
		}
		return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName)
	}

	dialogOptionId, err := sinfo.GetDialogOptionTag(parameters[0])
	if err != nil {
		return nil, err
	}

	convertedDialogId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(dialogOptionId.Id)
	if err != nil {
		return nil, err
	}

	retByteCode = append(retByteCode, cmd.commandId)
	retByteCode = append(retByteCode, convertedDialogId...)
	return retByteCode, nil
}

func (cmd *OptionCommand) isParametersCountCorrect(parameters []string) bool {
	return len(parameters) == 1
}