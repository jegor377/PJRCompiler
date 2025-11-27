package CommandFactory

import "errors"
import "pjre/Commands"

type CommandFactory struct {
	commands map[string]Command
}

func New() *CommandFactory {
	commands := map[string]Command {
		"bgColor": Commands.NewBgColorCommand(),
		"appName": Commands.NewAppNameCommand(),
		"wydt": Commands.NewWYDTCommand(),
		"text": Commands.NewTextCommand(),
		"color": Commands.NewColorCommand(),
		"option": Commands.NewOptionCommand(),
		"narrator": Commands.NewNarratorCommand(),
		"sprite": Commands.NewSpriteCommand(),
		"bgSong": Commands.NewBgSongCommand(),
	}
	return &CommandFactory{commands}
}

func (cmdFactory * CommandFactory) Get(key string) (Command, error) {
	if val, ok := cmdFactory.commands[key]; ok {
		return val, nil
	}
	errInfo := "Key [" + key + "] does not exist."
	return nil, errors.New(errInfo)
}