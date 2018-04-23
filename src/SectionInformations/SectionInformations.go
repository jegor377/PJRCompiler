package SectionInformations

import "errors"

type SectionInformations struct {
	SongTags map[string]uint32
	SpriteTags map[string]uint32
	DialogOptionTags map[string]DialogOption
	DialogTags map[string]uint32
}

func New() *SectionInformations {
	return &SectionInformations {
		make(map[string]uint32, 0),
		make(map[string]uint32, 0),
		make(map[string]DialogOption, 0),
		make(map[string]uint32, 0),
	}
}

func (sinfo *SectionInformations) AddSongTag(name string, id uint32) error {
	if _, ok := sinfo.SongTags[name]; !ok {
		sinfo.SongTags[name] = id
		return nil
	}
	return errors.New("Song override: " + name)
}

func (sinfo *SectionInformations) GetSongTag(name string) (uint32, error) {
	if v, ok := sinfo.SongTags[name]; ok {
		return v, nil
	}
	return 0, errors.New("No song tag with \"" + name + "\" name.")
}

func (sinfo *SectionInformations) AddSpriteTag(name string, id uint32) error {
	if _, ok := sinfo.SpriteTags[name]; !ok {
		sinfo.SpriteTags[name] = id
		return nil
	}
	return errors.New("Song override: " + name)
}

func (sinfo *SectionInformations) GetSpriteTag(name string) (uint32, error) {
	if v, ok := sinfo.SpriteTags[name]; ok {
		return v, nil
	}
	return 0, errors.New("No sprite tag with \"" + name + "\" name.")
}

func (sinfo *SectionInformations) AddDialogOptionTag(name string, target string, id uint32) error {
	if _, ok := sinfo.DialogOptionTags[name]; !ok {
		sinfo.DialogOptionTags[name] = DialogOption{id, target}
		return nil
	}
	return errors.New("Song override: " + name)
}

func (sinfo *SectionInformations) GetDialogOptionTag(name string) (DialogOption, error) {
	if v, ok := sinfo.DialogOptionTags[name]; ok {
		return v, nil
	}
	return DialogOption{0, ""}, errors.New("No dialog option tag with \"" + name + "\" name.")
}

func (sinfo *SectionInformations) AddDialogTag(name string, id uint32) error {
	if _, ok := sinfo.DialogTags[name]; !ok {
		sinfo.DialogTags[name] = id
		return nil
	}
	return errors.New("Song override: " + name)
}

func (sinfo *SectionInformations) GetDialogTag(name string) (uint32, error) {
	if v, ok := sinfo.DialogTags[name]; ok {
		return v, nil
	}
	return 0, errors.New("No dialog tag with \"" + name + "\" name.")
}