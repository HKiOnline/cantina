package character

import "encoding/json"

func Template() (string, error) {

	character := Character{}

	characterJson, err := json.Marshal(character)

	if err != nil {
		return "", err
	}

	return string(characterJson), nil
}
