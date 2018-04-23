package CommandFactory

import "SectionInformations"

type Command interface {
	GetByteCode([]string, string, string, *SectionInformations.SectionInformations) ([]byte, error)
}