package repository

import (
	"apps.hki.io/cantina/pkg/lib/character"
)

func NewRepository() character.Repository {

	repoPath := "db/characterCollection.pb.db"
	SetupProtobufCharacterRepository(repoPath)
	return ProtobufCharacterRepository{repositoryFilePath: repoPath}
}
