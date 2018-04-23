package PJRCompiler

import "errors"

func (compiler *PJRCompiler) compileSettings(labelSection []string) ([]byte, error) {
	retByteCode := make([]byte, 0)
	for _, v := range(labelSection) {
		if len(v) > 0 {
			splitedCommand := splitCommandAndParameters(v)
			section := "settings"
			label := "" // no labels in settings
			commandName := splitedCommand[0]
			if !isAvailableSettingsCommand(commandName) {
				return nil, errors.New("Not available settings command: " + commandName)
			}
			commandHandler, err := compiler.commandFactory.Get(commandName)
			if err != nil {
				return nil, err
			}
			byteCode, err := commandHandler.GetByteCode(splitedCommand[1:], section, label, compiler.sectionInfo)
			if err != nil {
				return nil, err
			}
			retByteCode = append(retByteCode, byteCode...)
		}
	}
	return retByteCode, nil
}


//TODO: popraw kod
func isAvailableSettingsCommand(commandName string) bool {
	availableSettingsCommands := []string{
		"bgColor",
		"appName",
		"wydt",
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