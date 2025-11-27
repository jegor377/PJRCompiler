package PJRCompiler

import "strings"

func (compiler *PJRCompiler) preprocessDialogOptions(labelSection []string) (map[string][]string, error) {
	labels, err := divideDialogOptionsByLabel(labelSection)
	// add tags to the stack
	var index uint32 = 0
	for n, _ := range(labels) {
		compiler.sectionInfo.AddDialogOptionTag(getDialogOptionName(n), getDialogOptionTarget(n), index)
		index++
	}
	return labels, err
}

// TODO: popraw kod xd
func divideDialogOptionsByLabel(labelSection []string) (map[string][]string, error) {
	dialogOptionsLabels := make(map[string][]string, 0)
	for i, v := range(labelSection) {
		if isDialogOptionLabel(v) {
			dialogOptionsLabels[v] = make([]string, 0)
			labelSectionSize := len(labelSection)
			if i + 1 < labelSectionSize {
				av := labelSection[i + 1]
				for j := i + 1; j < labelSectionSize && !isDialogOptionLabel(av); j++ {
					av = labelSection[j]
					if(!isDialogOptionLabel(av)) {
						dialogOptionsLabels[v] = append(dialogOptionsLabels[v], av)
					}
				}
			}
		}
	}
	return dialogOptionsLabels, nil
}

func isDialogOptionLabel(label string) bool {
	if len(label) > 0 {
		firstLetterIsMinus := []rune(label)[0] == rune('-')
		hasOnlyOneConnector := strings.Count(label, ":") == 1
		lastCharacterIsNotConnector := []rune(label)[len(label) - 1] != rune(':')
		return firstLetterIsMinus && hasOnlyOneConnector && lastCharacterIsNotConnector
	}
	return false
}

func getDialogOptionName(labelInformation string) string {
	return strings.Split(strings.TrimPrefix(labelInformation, "-"), ":")[0]
}

func getDialogOptionTarget(labelInformation string) string {
	return strings.Split(strings.TrimPrefix(labelInformation, "-"), ":")[1]
}