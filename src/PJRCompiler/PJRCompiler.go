package PJRCompiler

import "ScriptsManager"
import "CommandFactory"
import "SectionInformations"

type PJRCompiler struct {
	scriptsManager *ScriptsManager.ScriptsManager
	commandFactory *CommandFactory.CommandFactory
	sectionInfo *SectionInformations.SectionInformations
}

func New(sectionInfo *SectionInformations.SectionInformations) *PJRCompiler {
	return &PJRCompiler{
		nil,
		CommandFactory.New(),
		sectionInfo,
	}
}

func (compiler *PJRCompiler) Load(scriptsManager *ScriptsManager.ScriptsManager) {
	compiler.scriptsManager = scriptsManager
}