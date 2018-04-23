package ScriptsManager

import "io/ioutil"
import "os"
import "strings"

func (sm *ScriptsManager) LoadScripts(folderPath string) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}
	sm.clearScripts()
	err = sm.loadScriptFiles(folderPath, files)
	if err != nil {
		return err
	}
	return nil
}

func (sm *ScriptsManager) loadScriptFiles(folderPath string, files []os.FileInfo) error {
	for _, file := range(files) {
		if !file.IsDir() {
			tmpPath := folderPath + "\\" + file.Name()
			err := sm.loadScriptFile(tmpPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (sm *ScriptsManager) loadScriptFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	sm.Scripts = append( sm.Scripts, removeShitFromScript(string(file)) )
	return nil
}

// I had no idea how to name that method.
func removeShitFromScript(script string) []string {
	lines := strings.Split(strings.TrimSpace(script), "\n")
	for k, v := range(lines) {
		lines[k] = strings.TrimSpace(v)
	}
	return lines
}