package character

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"apps.hki.io/cantina/pkg/lib/calendar"
)

func (character *Character) ShowJson() (string, error) {

	characterJson, err := json.Marshal(character)

	if err != nil {
		return "", err
	}

	return string(characterJson), nil

}

func (character *Character) ShowCondensedText() (string, error) {

	tmpl, err := template.New("condensed.txt").Funcs(template.FuncMap{
		"title":                 strings.Title,
		"formatYear":            formatYear,
		"formatForceUsesStatus": formatForceUsesStatus,
	}).ParseFiles(
		"pkg/template/character/basic_information.txt",
		"pkg/template/character/specializations.txt",
		"pkg/template/character/vitals.txt",
		"pkg/template/character/characteristics.txt",
		"pkg/template/character/skills.txt",
		"pkg/template/character/talents.txt",
		"pkg/template/character/force.txt",
		"pkg/template/character/equipment.txt",
		"pkg/template/character/condensed.txt")

	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, character)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func formatYear(year int32) string {
	return calendar.BattleOfYavinToHumanReadable(int(year))
}
