package sheet

import (
	"os"

	"apps.hki.io/cantina/pkg/lib/character"
	"github.com/spf13/cobra"
)

var CharacterSheetCmd = &cobra.Command{
	Use:   "sheet",
	Short: "The character sheet command generates a Star Wars Roleplaying character sheet in JSON-format.",
	Long:  `The character sheet command generates a Star Wars Roleplaying character sheet in JSON-format.`,
	Run: func(cmd *cobra.Command, args []string) {

		s, _ := character.Template()
		os.Stdout.WriteString(s)
	},
}

func init() {
}
