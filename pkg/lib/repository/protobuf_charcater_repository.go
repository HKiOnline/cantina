package repository

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"apps.hki.io/cantina/pkg/lib/character"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"
	"google.golang.org/protobuf/proto"
)

type ProtobufCharacterRepository struct {
	repositoryFilePath string
}

func (pcr ProtobufCharacterRepository) Insert(ctx context.Context, inCharacter character.Character) (character.Character, error) {

	collection, err := getCollection(pcr.repositoryFilePath)

	if err != nil {
		return character.Character{}, err
	}

	// Create an id for the character if there isn't one
	if inCharacter.Id == "" {
		inCharacter.Id = uuid.New().String()[0:8]
	}

	// Check for duplicates
	for _, stored := range collection.Characters {

		if stored.Name == inCharacter.Name {
			return character.Character{}, errors.New("character with the same name already exists")
		}

		if stored.Id == inCharacter.Id {
			return character.Character{}, errors.New("character with the same id already exists")
		}
	}

	// Add the new charcter to the collection
	collection.Characters = append(collection.Characters, &inCharacter)

	err = saveCollection(pcr.repositoryFilePath, &collection)
	if err != nil {
		return character.Character{}, err
	}

	// Return the character and zero errors
	return inCharacter, nil

}

func (pcr ProtobufCharacterRepository) Get(ctx context.Context, id string) (character.Character, error) {

	collection, err := getCollection(pcr.repositoryFilePath)

	if err != nil {
		return character.Character{}, err
	}

	// Check for duplicates
	for index, stored := range collection.Characters {
		if stored.Id == id {
			return *collection.Characters[index], nil
		}
	}

	return character.Character{}, errors.New("character with the given id was not found")
}

func (pcr ProtobufCharacterRepository) GetWithName(ctx context.Context, name string) (character.Character, error) {

	collection, err := getCollection(pcr.repositoryFilePath)

	if err != nil {
		return character.Character{}, err
	}

	// Check for duplicates
	for index, stored := range collection.Characters {
		if stored.Name == strings.TrimSpace(name) {
			return *collection.Characters[index], nil
		}
	}

	return character.Character{}, errors.New("character with the given name was not found")
}

func (pcr ProtobufCharacterRepository) Delete(ctx context.Context, id string) (character.Character, error) {

	collection, err := getCollection(pcr.repositoryFilePath)

	if err != nil {
		return character.Character{}, err
	}

	idFound := false
	oldCharacter := character.Character{}

	for index, stored := range collection.Characters {

		if stored.Id == id {
			fmt.Printf("Found %s at index %d\n", id, index)
			idFound = true
			oldCharacter = *stored

			if len(collection.Characters) == 1 {
				collection = character.Collection{}
			} else {
				collection.Characters = slices.Delete(collection.Characters, index, index+1)
			}

		}
	}

	if idFound {

		err = saveCollection(pcr.repositoryFilePath, &collection)
		if err != nil {
			return oldCharacter, err
		} else {
			return oldCharacter, nil
		}

	} else {
		return character.Character{}, errors.New("character with the given id was not found")
	}

}
func (pcr ProtobufCharacterRepository) List(ctx context.Context) ([]character.Character, error) {

	collection, err := getCollection(pcr.repositoryFilePath)

	if err != nil {
		return []character.Character{}, err
	}

	characters := []character.Character{}
	for _, cPointers := range collection.Characters {
		characters = append(characters, *cPointers)
	}

	return characters, nil
}

func getCollection(filepath string) (character.Collection, error) {

	// Open the character repository DB file
	in, err := os.ReadFile(filepath)
	if err != nil {
		return character.Collection{}, err
	}

	// Create Go-structs out of the raw binary information
	cc := character.Collection{}
	if err := proto.Unmarshal(in, &cc); err != nil {
		return character.Collection{}, err
	}

	return cc, nil
}

func saveCollection(filepath string, collection *character.Collection) error {

	// Transform the appended collection back to protobuf binary format
	out, err := proto.Marshal(collection)
	if err != nil {
		return err
	}

	// Save the new file
	if err := os.WriteFile(filepath, out, 0644); err != nil {
		return err
	}

	// Return zero errors hopefully
	return nil
}

func SetupProtobufCharacterRepository(repoPath string) {

	if _, err := os.Stat(repoPath); errors.Is(err, os.ErrNotExist) {
		collection := character.Collection{}
		saveCollection(repoPath, &collection)
	}
}
