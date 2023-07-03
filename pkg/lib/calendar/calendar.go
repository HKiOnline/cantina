package calendar

import "fmt"

const (
	BeforeBattleOfYavinStr string = "BBY"
	AfterBattleOfYavinStr  string = "ABY"
)

func epochToHumanReadable(year int, before string, after string) string {

	if year < 0 {
		year = -1 * year
		return fmt.Sprintf("%d %s", year, before)
	} else {
		return fmt.Sprintf("%d %s", year, after)
	}

}

func BattleOfYavinToHumanReadable(year int) string {
	return epochToHumanReadable(year, BeforeBattleOfYavinStr, AfterBattleOfYavinStr)
}
