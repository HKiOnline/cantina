package character

import (
	"apps.hki.io/cantina/pkg/cmd/character/add"
	"apps.hki.io/cantina/pkg/cmd/character/delete"
	"apps.hki.io/cantina/pkg/cmd/character/list"
	"apps.hki.io/cantina/pkg/cmd/character/read"
	"apps.hki.io/cantina/pkg/cmd/character/sheet"
	"apps.hki.io/cantina/pkg/cmd/character/show"
	"github.com/spf13/cobra"
)

var CharacterCmd = &cobra.Command{
	Use:   "character",
	Short: "The character command helps handling Star Wars Roleplaying characters.",
	Long: `The character command helps handling Star Wars Roleplaying characters.
To use the command, you need to add an additional subcommand.`,
}

func addSubcommands() {

	CharacterCmd.AddCommand(add.AddCharacterCmd)
	CharacterCmd.AddCommand(list.ListCharacterCmd)
	CharacterCmd.AddCommand(delete.DeleteCharacterCmd)

	CharacterCmd.AddCommand(sheet.CharacterSheetCmd)
	CharacterCmd.AddCommand(show.ShowCharacterCmd)
	CharacterCmd.AddCommand(read.ReadCharacterCmd)
}

func init() {
	addSubcommands()
}
