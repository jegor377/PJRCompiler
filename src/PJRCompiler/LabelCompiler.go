package PJRCompiler

import "strings"

func compileLabel(name string, path string) []byte {
	retByteCode := make([]byte, 0)
	// Convert string to byte array
	convertedPath := []byte(path)
	convertedPath = append(convertedPath, byte(0)) // adds zero character at the end
	// Assembly correct bytecode
	retByteCode = append(retByteCode, convertedPath...)
	return retByteCode
}

func getLabelNameAndPath(labelInformation string) (string, string) {
	labelName := getLabelName(labelInformation)
	labelPath := getLabelPath(labelInformation)
	return labelName, labelPath
}

func getLabelName(labelInformation string) string {
	return strings.Split(strings.TrimPrefix(labelInformation, "-"), ";")[0]
}

func getLabelPath(labelInformation string) string {
	return strings.Split(strings.TrimPrefix(labelInformation, "-"), ";")[1]
}

func isCorrectLabelInformation(labelInformation string) bool {
	firstCharacterIsOk := []rune(labelInformation)[0] == rune('-')
	connectorIsNotTheLastCharacter := []rune(labelInformation)[len(labelInformation)-1] != rune(';')
	thereIsOnlyOneConnector := strings.Count(labelInformation, ";") == 1
	//thereIsOnlyOneFirstCharacter := strings.Count(labelInformation, "-") == 1
	return firstCharacterIsOk && connectorIsNotTheLastCharacter && thereIsOnlyOneConnector /*&& thereIsOnlyOneFirstCharacter*/
}