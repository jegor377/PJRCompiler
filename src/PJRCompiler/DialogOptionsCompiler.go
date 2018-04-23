package PJRCompiler

import "errors"
import "ConvertingTools"

// TODO: Zdecydowanie popraw kod
func (compiler *PJRCompiler) compileDialogOptions(processedLabels map[string][]string) ([]byte, error) {
	retByteCode := make([]byte, 0)
	for n, e := range(processedLabels) {
		dialogOptionName := getDialogOptionName(n)
		optionTag, err := compiler.sectionInfo.GetDialogOptionTag(dialogOptionName)
		if err != nil {
			return nil, err
		}
		dialogOptionId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(optionTag.Id)
		if err != nil {
			return nil, err
		}

		targetId, err := compiler.sectionInfo.GetDialogTag(optionTag.Target)
		if err != nil {
			return nil, err
		}
		convertedTargetId, err := ConvertingTools.ConvertUint32ToBytesLittleEndian(targetId)
		if err != nil {
			return nil, err
		}

		labelByteCode := make([]byte, 0)

		retByteCode = append(retByteCode, dialogOptionId...) // adds dialog option id to the byte code stack ...[id]...
		retByteCode = append(retByteCode, convertedTargetId...) // adds target dialog id to the byte code stack ...[id][target]...

		// collect byte code from commands
		for _, v := range(e) {
			if len(v) > 0 {
				splitedCommand := splitCommandAndParameters(v)
				section := "dialogOptions"
				commandName := splitedCommand[0]
				if !isAvailableDialogOptionsCommand(commandName) {
					return nil, errors.New("Not available dialog option command: " + commandName)
				}
				commandHandler, err := compiler.commandFactory.Get(commandName)
				if err != nil {
					return nil, err
				}
				byteCode, err := commandHandler.GetByteCode(splitedCommand[1:], section, dialogOptionName, compiler.sectionInfo)
				if err != nil {
					return nil, err
				}
				labelByteCode = append(labelByteCode, byteCode...)
			}
		}

		labelSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32( len(labelByteCode) ) ) // get byte code size in bytes
		if err != nil {
			return nil, err
		}


		retByteCode = append(retByteCode, labelSize...) // add byte code size to the byte code stack ...[id][target][size]...
		retByteCode = append(retByteCode, labelByteCode...) // add byte code to the byte code stack ...[id][target][size][bytecode]...
	}
	return retByteCode, nil
}


//TODO: popraw kod
func isAvailableDialogOptionsCommand(commandName string) bool {
	availableSettingsCommands := []string{
		"bgColor",
		"appName",
		"wydt",
		"text",
		"color",
		"bgSong",
	}
	for _, v := range(availableSettingsCommands) {
		if commandName == v {
			return true
		}
	}
	return false
}