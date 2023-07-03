package character

func (character *Character) HasCareer() bool {
	if character.Career != nil {
		return true
	} else {
		return false
	}
}

func (character *Character) HasSpecializations() bool {
	if character.Career != nil && character.Career.Specializations != nil {
		return true
	} else {
		return false
	}
}
