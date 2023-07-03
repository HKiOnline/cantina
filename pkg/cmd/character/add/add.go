package add

import (
	"context"
	"fmt"
	"log"
	"os"

	"apps.hki.io/cantina/pkg/lib/character"
	"apps.hki.io/cantina/pkg/lib/repository"
	"github.com/spf13/cobra"
)

var AddCharacterCmd = &cobra.Command{
	Use:   "add",
	Short: "The character add commands adds a Star Wars Roleplaying character's information to character database.",
	Long:  `The character add commands adds a Star Wars Roleplaying character's information to character database.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 1 {
			log.Fatal("Incorret arguments. Exactly one argument, file path to character sheet, is needed.")
		}

		// 1. Get the character sheet file as file path from the commandline
		filepath := args[0]

		// 2. Open the file and parse the character sheet
		characterFromFile, err := character.ReadFromFile(filepath)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(characterFromFile.Characteristics)

		// 3. Get a new character repository and
		// store the character's information
		characterRepository := repository.NewRepository()
		addedCharacter, err := characterRepository.Insert(context.Background(), characterFromFile)

		if err != nil {
			log.Fatal(err)
		}

		// 4. Print out information such as the character's unique identifier
		os.Stdout.WriteString(fmt.Sprintf("Character %s added with id %s", addedCharacter.Name, addedCharacter.Id))
	},
}

func init() {
}
