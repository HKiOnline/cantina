package delete

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"apps.hki.io/cantina/pkg/lib/repository"
	"github.com/spf13/cobra"
)

var DeleteCharacterCmd = &cobra.Command{
	Use:   "delete",
	Short: "The character delete commands removes a Star Wars Roleplaying character's information from the character database.",
	Long:  `The character delete commands removes a Star Wars Roleplaying character's information from the character database.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			log.Fatal("Incorret arguments. Exactly one argument, character id, is needed.")
		}

		// 1. Get the character sheet file as file path from the commandline
		id := strings.TrimSpace(args[0])

		// 2. Get the character repository and
		// delete the character
		characterRepository := repository.NewRepository()
		deletedCharacter, err := characterRepository.Delete(context.Background(), id)

		if err != nil {
			log.Fatal(err)
		}

		// 4. Print out information such as the character's unique identifier
		os.Stdout.WriteString(fmt.Sprintf("Character %s with id %s was deleted", deletedCharacter.Name, deletedCharacter.Id))
	},
}

func init() {
}
