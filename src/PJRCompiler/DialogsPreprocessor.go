package PJRCompiler

import "strings"
import "errors"

func (compiler *PJRCompiler) preprocessDialogs(labelSection []string) (map[string][]string, error) {
	labels, err := divideDialogsByLabel(labelSection)
	if err != nil {
		return nil, err
	}
	// add tags to the stack
	var index uint32 = 1
	compiler.sectionInfo.AddDialogTag("main", 0)
	for n, _ := range(labels) {
		name := getDialogName(n)
		if name != "main" {
			compiler.sectionInfo.AddDialogTag(name, index)
			index++
		}
	}
	return labels, nil
}

// TODO: popraw kod xd
func divideDialogsByLabel(labelSection []string) (map[string][]string, error) {
	dialogsLabels := make(map[string][]string, 0)
	for i, v := range(labelSection) {
		if isDialogLabel(v) {

			if _, ok := dialogsLabels[v]; ok {
				return nil, errors.New("Two main labels declared. There can be only 1.")
			}
			dialogsLabels[v] = make([]string, 0)

			labelSectionSize := len(labelSection)
			if i + 1 < labelSectionSize {
				av := labelSection[i + 1]
				for j := i + 1; j < labelSectionSize && !isDialogLabel(av); j++ {
					av = labelSection[j]
					if(!isDialogLabel(av)) {
						dialogsLabels[v] = append(dialogsLabels[v], av)
					}
				}
			}
		}
	}
	return dialogsLabels, nil
}

func isDialogLabel(label string) bool {
	if len(label) > 0 {
		firstLetterIsMinus := []rune(label)[0] == rune('-')
		lastCharacterIsNotMinus := []rune(label)[len(label) - 1] != rune('-')
		return firstLetterIsMinus && lastCharacterIsNotMinus
	}
	return false
}

func getDialogName(labelInformation string) string {
	return strings.TrimPrefix(labelInformation, "-")
}