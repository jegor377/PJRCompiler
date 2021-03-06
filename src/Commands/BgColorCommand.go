package Commands

import "SectionInformations"
import "strconv"
import "errors"
import "strings"
import "ConvertingTools"

// bgColor NUMBER:0-1
type BgColorCommand struct {
	commandId uint8
	name string
}

func NewBgColorCommand() *BgColorCommand {
	return &BgColorCommand{0, "bgColor"}
}

// [command id:1 byte][parameter:1 byte]
func (cmd *BgColorCommand) GetByteCode(parameters []string, sectionName string, labelName string, sinfo *SectionInformations.SectionInformations) ([]byte, error) {
	retByteCode := make([]byte, 0)
	if !cmd.isParametersCountCorrect(parameters) {
		if len(labelName) > 0 {
			return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName + ", label: " + labelName)
		}
		return nil, errors.New("Parameters count is incorrect: " + cmd.name + " \"" + strings.Join(parameters, "\" \"") + "\", section name: " + sectionName)
	}

	colorId, err := strconv.Atoi(parameters[0]) // It takes only one uint32 parameter
	if err != nil {
		return nil, err
	}
	if !cmd.isColorIdParameterCorrect(colorId) {
		if len(labelName) > 0 {
			return nil, errors.New("Parameter value is incorrect: " + cmd.name + " \"" + parameters[0] + "\", section name: " + sectionName + ", label: " + labelName + ". It can be only 0 or 1.")
		} 
		return nil, errors.New("Parameter value is incorrect: " + cmd.name + " \"" + parameters[0] + "\", section name: " + sectionName + ". It can be only 0 - 16.")
	}
	convertedColorId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(uint32(colorId))
	if err != nil {
		return nil, err
	}

	retByteCode = append(retByteCode, cmd.commandId)
	retByteCode = append(retByteCode, convertedColorId[0]) // It needs only 0 - 16 value, so I pass only one byte to make it smaller.
	return retByteCode, nil
}

func (cmd *BgColorCommand) isParametersCountCorrect(parameters []string) bool {
	return len(parameters) == 1
}

func (cmd *BgColorCommand) isColorIdParameterCorrect(val int) bool {
	return val >= 0 && val <= 16
}