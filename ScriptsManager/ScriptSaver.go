package ScriptsManager

import "os"

func (sm *ScriptsManager) SaveScript(filePath string, byteCode []byte) error {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	f.Write(byteCode)
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}