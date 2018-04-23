package PJRCompiler

import "errors"

func (compiler *PJRCompiler) compileSongs(labelSection []string) ([]byte, error) {
	retByteCode := make([]byte, 0)
	for _, v := range(labelSection) {
		if len(v) > 0 {
			if isCorrectLabelInformation(v) {
				labelName, labelPath := getLabelNameAndPath(v)

				if len(labelName) > 0 && len(labelPath) > 0 {
					id := uint32(len(compiler.sectionInfo.SongTags))
					err := compiler.sectionInfo.AddSongTag(labelName, id)
					if err != nil {
						return nil, err
					}

					compiledLabel := compileLabel(labelName, labelPath)
					retByteCode = append(retByteCode, compiledLabel...)
				} else {
					return nil, errors.New("Label name or path is emtpy: " + v)
				}
			} else {
				return nil, errors.New("Label information is incorrect: " + v)
			}
		}
	}
	return retByteCode, nil
}