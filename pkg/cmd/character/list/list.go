package list

import (
	"context"
	"log"
	"os"

	"apps.hki.io/cantina/pkg/lib/repository"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var ListCharacterCmd = &cobra.Command{
	Use:   "list",
	Short: "The character list commands lists stored Star Wars Roleplaying characters.",
	Long:  `The character list commands lists stored Star Wars Roleplaying characters.`,
	Run: func(cmd *cobra.Command, args []string) {

		characterRepository := repository.NewRepository()
		characters, err := characterRepository.List(context.Background())

		if err != nil {
			log.Fatal(err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"id", "Name", "Type", "Career", "Specializations"})

		for _, character := range characters {

			careerName := "-"
			specializations := "-"

			if character.HasCareer() {
				careerName = character.Career.Name
			}

			if character.HasSpecializations() {
				specializations = character.ListSpecializations()
			}

			t.AppendRow([]interface{}{character.Id, character.Name, character.GetType(), careerName, specializations})
			t.AppendSeparator()
		}

		t.Render()

	},
}

func init() {
}
