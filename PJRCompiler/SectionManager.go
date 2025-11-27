package PJRCompiler

func (compiler *PJRCompiler) getAllEntitiesFromSection(sectionName string) []string {
	sectionEntities := make([]string, 0)
	for _, v := range(compiler.scriptsManager.Scripts) {
		sectionEntities = append(sectionEntities, searchForSection(v, sectionName)...)
	}
	return sectionEntities
}

func searchForSection(script []string, sectionName string) []string {
	ret := make([]string, 0)
	for i, v := range(script) {
		if isSectionName(v) {
			if foundSection(v, sectionName) {
				// TODO: Poprawić kod
				scriptLen := len(script)
				if i + 1 < scriptLen {
					av := script[i + 1]
					for j := i + 1; !isSectionName(av) && j<scriptLen; j++ {
						av = script[j]
						if !isSectionName(av) && len(av) > 0 {
							ret = append(ret, av)
						}
					}
				}
			}
		}
	}
	return ret
}

func foundSection(tmp string, sectionName string) bool {
	tmpSectionName := ">" + sectionName
	return tmp == tmpSectionName
}

func isSectionName(name string) bool {
	if(len(name) > 0) {
		return []rune(name)[0] == '>'
	}
	return false
}