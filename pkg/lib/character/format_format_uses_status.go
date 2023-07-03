package character

func formatForceUsesStatus(forceRank int32, morality int32) string {

	if forceRank > 0 {
		switch {
		case morality < 30:
			return "Dark Side Force User"
		case morality >= 30:
			return "Light Side Force User"
		case morality > 70:
			return "Paragon Light Side Force User"
		default:
			return "Force User"
		}
	} else {
		return ""
	}
}
