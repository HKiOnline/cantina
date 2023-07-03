package character

import "strings"

func (character *Character) ListSpecializations() string {

	if character.HasSpecializations() {
		specs := []string{}
		for _, spec := range character.Career.Specializations {
			specs = append(specs, spec.Name)
		}

		return strings.Join(specs, ", ")
	} else {
		return ""
	}

}
