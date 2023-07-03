package show

import (
	"context"
	"log"
	"os"
	"strings"

	"apps.hki.io/cantina/pkg/lib/character"
	"apps.hki.io/cantina/pkg/lib/repository"
	"github.com/spf13/cobra"
)

var ShowCharacterCmd = &cobra.Command{
	Use:   "show",
	Short: "The the show command outputs a Star Wars Roleplaying character's information.",
	Long:  `The the show command outputs a Star Wars Roleplaying character's information.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			cmd.Help()
			os.Exit(1)
		}

		characterRepository := repository.NewRepository()

		nameFlag, err := cmd.Flags().GetBool("name")
		if err != nil {
			log.Fatal(err)
		}

		showCharacter := character.Character{}

		if nameFlag {
			searchKey := strings.TrimSpace(strings.Join(args, " "))
			showCharacter, err = characterRepository.GetWithName(context.Background(), searchKey)

			if err != nil {
				log.Fatal(err)
			}

		} else {
			searchKey := strings.TrimSpace(args[0])
			showCharacter, err = characterRepository.Get(context.Background(), searchKey)

			if err != nil {
				log.Fatal(err)
			}
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatal(err)
		}

		var characterStr string

		if jsonFlag {
			characterStr, err = showCharacter.ShowJson()
		} else {
			characterStr, err = showCharacter.ShowCondensedText()
		}

		if err != nil {
			log.Fatal(err)
		}

		os.Stdout.WriteString(characterStr)
	},
}

func init() {

	//character.CharacterCmd.AddCommand(sheetCmd)
	ShowCharacterCmd.Flags().BoolP("json", "j", false, "displays a un-formated parsed version of the character data in JSON")
	ShowCharacterCmd.Flags().BoolP("name", "n", false, "find by name")
}
