package ScriptsManager

type ScriptsManager struct {
	Scripts [][]string
}

func New() *ScriptsManager {
	return &ScriptsManager{make([][]string, 0)}
}

func (sm *ScriptsManager) clearScripts() {
	sm.Scripts = make([][]string, 0)
}