package PJRCompiler

import "errors"
import "pjre/ConvertingTools"

func (compiler *PJRCompiler) compileDialogs(processedLabels map[string][]string) ([]byte, error) {
	retByteCode := make([]byte, 0)
	for n, e := range(processedLabels) {
		dialogName := getDialogName(n)
		dialogId, err := compiler.sectionInfo.GetDialogTag(dialogName)
		if err != nil {
			return nil, err
		}
		convertedDialogId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(dialogId) // get dialog id
		if err != nil {
			return nil, err
		}

		labelByteCode := make([]byte, 0) // create tmp stack for byte code

		retByteCode = append(retByteCode, convertedDialogId...) // add dialog id to the byte code stack (main) ...[id]...
		if isEndingDialog(dialogName) {
			retByteCode = append(retByteCode, byte(1))
		} else {
			retByteCode = append(retByteCode, byte(0))
		}

		// create byte code in tmp
		for _, v := range(e) {
			if len(v) > 0 {
				splitedCommand := splitCommandAndParameters(v)
				section := "dialogs"
				commandName := splitedCommand[0]
				if !isAvailableDialogsCommand(commandName) {
					return nil, errors.New("Not available dialog command: " + commandName)
				}
				commandHandler, err := compiler.commandFactory.Get(commandName)
				if err != nil {
					return nil, err
				}
				byteCode, err := commandHandler.GetByteCode(splitedCommand[1:], section, dialogName, compiler.sectionInfo)
				if err != nil {
					return nil, err
				}
				labelByteCode = append(labelByteCode, byteCode...)
			}
		}

		labelSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32( len(labelByteCode) ) ) // get tmp byte code size
		if err != nil {
			return nil, err
		}

		retByteCode = append(retByteCode, labelSize...) // add tmp byte code size to the byte code stack ...[id][size]...
		retByteCode = append(retByteCode, labelByteCode...) // add tmp byte code to the byte code stack ...[id][size][byte code]...
	}
	return retByteCode, nil
}


//TODO: popraw kod
func isAvailableDialogsCommand(commandName string) bool {
	availableSettingsCommands := []string{
		"bgColor",
		"appName",
		"wydt",
		"text",
		"color",
		"option",
		"narrator",
		"sprite",
		"bgSongs",
	}
	for _, v := range(availableSettingsCommands) {
		if commandName == v {
			return true
		}
	}
	return false
}

func isEndingDialog(dialogName string) bool {
	return []rune(dialogName)[0] == rune('!')
}