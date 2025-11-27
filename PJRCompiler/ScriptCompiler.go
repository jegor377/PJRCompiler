package PJRCompiler

import "errors"
import "pjre/ConvertingTools"

// TODO: ZDECYDOWANIE popraw kod
func (compiler *PJRCompiler) Compile() ([]byte, error) {
	if len(compiler.scriptsManager.Scripts) == 0 {
		return nil, errors.New("Cannot to compile empty list of scripts.")
	}

	retByteCode := make([]byte, 0)
	retByteCode = append(retByteCode, []byte("PJR")...)
	
	spritesByteCode, err := compiler.getSpritesSectionByteCode()
	if err != nil {
		return nil, err
	}
	spritesSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32(len(spritesByteCode)) )
	if err != nil {
		return nil, err
	}

	songsByteCode, err := compiler.getSongsSectionByteCode()
	if err != nil {
		return nil, err
	}
	songsSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32(len(songsByteCode)) )
	if err != nil {
		return nil, err
	}

	settingsByteCode, err := compiler.getSettingsSectionByteCode()
	if err != nil {
		return nil, err
	}
	settingsSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32(len(settingsByteCode)) )
	if err != nil {
		return nil, err
	}

	dialogOptionsProcessedCode, err := compiler.getDialogOptionsProcessedCode()
	if err != nil {
		return nil, err
	}
	dialogsProcessedCode, err := compiler.getDialogsProcessedCode()
	if err != nil {
		return nil, err
	}

	dialogOptionsByteCode, err := compiler.compileDialogOptions(dialogOptionsProcessedCode)
	if err != nil {
		return nil, err
	}
	dialogOptionsSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32(len(dialogOptionsByteCode)) )
	if err != nil {
		return nil, err
	}

	dialogsByteCode, err := compiler.compileDialogs(dialogsProcessedCode)
	if err != nil {
		return nil, err
	}
	dialogsSize, err := ConvertingTools.ConvertUint32ToBytesLittleEndian( uint32(len(dialogsByteCode)) )
	if err != nil {
		return nil, err
	}

	retByteCode = append(retByteCode, spritesSize...)
	retByteCode = append(retByteCode, songsSize...)
	retByteCode = append(retByteCode, settingsSize...)
	retByteCode = append(retByteCode, dialogOptionsSize...)
	retByteCode = append(retByteCode, dialogsSize...)

	if len(spritesByteCode) > 0 {
		retByteCode = append(retByteCode, spritesByteCode...)
	}
	if len(songsByteCode) > 0 {
		retByteCode = append(retByteCode, songsByteCode...)
	}
	if len(settingsByteCode) > 0 {
		retByteCode = append(retByteCode, settingsByteCode...)
	}
	if len(dialogOptionsByteCode) > 0 {
		retByteCode = append(retByteCode, dialogOptionsByteCode...)
	}
	if len(dialogsByteCode) > 0 {
		retByteCode = append(retByteCode, dialogsByteCode...)
	}
	
	return retByteCode, nil
}

func (compiler *PJRCompiler) getSongsSectionByteCode() ([]byte, error) {
	songsEntities := compiler.getAllEntitiesFromSection("songs")
	if len(songsEntities) > 0 {
		songsByteCode, err := compiler.compileSongs(songsEntities)
		return songsByteCode, err
	}
	return nil, nil
}

func (compiler *PJRCompiler) getSpritesSectionByteCode() ([]byte, error) {
	spritesEntities := compiler.getAllEntitiesFromSection("sprites")
	if len(spritesEntities) > 0 {
		spriteByteCode, err := compiler.compileSprites(spritesEntities)
		return spriteByteCode, err
	}
	return nil, nil
}

func (compiler *PJRCompiler) getSettingsSectionByteCode() ([]byte, error) {
	settingsEntities := compiler.getAllEntitiesFromSection("settings")
	if len(settingsEntities) > 0 {
		settingsByteCode, err := compiler.compileSettings(settingsEntities)
		return settingsByteCode, err
	}
	return nil, nil
}

func (compiler *PJRCompiler) getDialogOptionsProcessedCode() (map[string][]string, error) {
	dialogOptionsEntities := compiler.getAllEntitiesFromSection("dialogOptions")
	if len(dialogOptionsEntities) > 0 {
		dialogOptionsProcessedCode, err := compiler.preprocessDialogOptions(dialogOptionsEntities)
		return dialogOptionsProcessedCode, err
	}
	return nil, nil
}

func (compiler *PJRCompiler) getDialogsProcessedCode() (map[string][]string, error) {
	dialogsEntities := compiler.getAllEntitiesFromSection("dialogs")
	if len(dialogsEntities) > 0 {
		dialogsProcessedCode, err := compiler.preprocessDialogs(dialogsEntities)
		return dialogsProcessedCode, err
	}
	return nil, nil
}