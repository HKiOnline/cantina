package character

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadFromFile(filepath string) (Character, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return Character{}, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)

	if err != nil {
		return Character{}, err
	}

	var character Character
	err = json.Unmarshal(bytes, &character)

	if err != nil {
		return Character{}, err
	} else {
		return character, nil
	}

}
