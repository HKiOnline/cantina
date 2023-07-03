package cmd

import (
	"fmt"
	"os"

	"apps.hki.io/cantina/pkg/cmd/character"
	"github.com/spf13/cobra"
)

var CantinaCmd = &cobra.Command{
	Use:   "cantina",
	Short: "The cantina command helps handling Star Wars Roleplaying data.",
	Long: `The cantina command helps handling Star Wars Roleplaying data.
To use the command, you need to add an additional subcommand.`,
}

func Execute() {
	if err := CantinaCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func addSubcommands() {
	CantinaCmd.AddCommand(character.CharacterCmd)
}

func init() {
	addSubcommands()
}
