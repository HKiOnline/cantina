package read

import (
	"log"
	"os"

	"apps.hki.io/cantina/pkg/lib/character"
	"github.com/spf13/cobra"
)

var ReadCharacterCmd = &cobra.Command{
	Use:   "read",
	Short: "The the read command reads a Star Wars Roleplaying character's information from JSON data and outputs it in other formats.",
	Long:  `The the read command reads a Star Wars Roleplaying character's information from JSON data and outputs it in other formats.`,
	Run: func(cmd *cobra.Command, args []string) {

		filepath := args[0]

		characterSheet, err := character.ReadFromFile(filepath)

		if err != nil {
			log.Fatal(err)
		}

		jsonFlag, err := cmd.Flags().GetBool("json")
		if err != nil {
			log.Fatal(err)
		}

		var characterStr string

		if jsonFlag == true {
			characterStr, err = characterSheet.ShowJson()
		} else {
			characterStr, err = characterSheet.ShowCondensedText()
		}

		if err != nil {
			log.Fatal(err)
		}

		os.Stdout.WriteString(characterStr)
	},
}

func init() {

	//character.CharacterCmd.AddCommand(sheetCmd)
	ReadCharacterCmd.Flags().BoolP("json", "j", false, "displays a un-formated parsed version of the character data in JSON")

}
