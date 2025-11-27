package CommandFactory

import "pjre/SectionInformations"

type Command interface {
	GetByteCode([]string, string, string, *SectionInformations.SectionInformations) ([]byte, error)
}