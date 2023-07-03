package character

func (charcter *Character) GetType() string {

	if charcter.Npc == nil {
		return "PC"
	} else {
		return charcter.Npc.Rank
	}

}
