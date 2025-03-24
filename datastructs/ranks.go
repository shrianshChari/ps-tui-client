package datastructs

type group struct {
	Name   string
	Symbol string
	Type   string
	Order  int
}

// Group definitions defined in:
// https://github.com/smogon/pokemon-showdown-client/blob/master/play.pokemonshowdown.com/src/client-main.ts#L342
func roomOwner() group {
	return group{
		Name:   "Room Owner",
		Symbol: "#",
		Type:   "leadership",
		Order:  101,
	}
}

func administrator() group {
	return group{
		Name:   "Administrator",
		Symbol: "~",
		Type:   "leadership",
		Order:  102,
	}
}

func administrator2() group {
	return group{
		Name:   "Administrator",
		Symbol: "&",
		Type:   "leadership",
		Order:  103,
	}
}

func host() group {
	return group{
		Name:   "Host",
		Symbol: "\u2605",
		Type:   "staff",
		Order:  104,
	}
}

func moderator() group {
	return group{
		Name:   "Moderator",
		Symbol: "@",
		Type:   "staff",
		Order:  105,
	}
}

func driver() group {
	return group{
		Name:   "Driver",
		Symbol: "%",
		Type:   "staff",
		Order:  106,
	}
}

func bot() group {
	return group{
		Name:   "Bot",
		Symbol: "*",
		Type:   "normal",
		Order:  109,
	}
}

func player() group {
	return group{
		Name:   "Player",
		Symbol: "\u2606",
		Type:   "normal",
		Order:  110,
	}
}

func voice() group {
	return group{
		Name:   "Voice",
		Symbol: "+",
		Type:   "normal",
		Order:  200,
	}
}

func normal() group {
	return group{
		Name:   "",
		Symbol: " ",
		Type:   "normal",
		Order:  201,
	}
}

func muted() group {
	return group{
		Name:   "Muted",
		Symbol: "!",
		Type:   "punishment",
		Order:  301,
	}
}

func namelocked() group {
	return group{
		Name:   "Namelocked",
		Symbol: "\u2716",
		Type:   "punishment",
		Order:  302,
	}
}

func locked() group {
	return group{
		Name:   "Locked",
		Symbol: "\u203d",
		Type:   "punishment",
		Order:  303,
	}
}

func GetGroup(symbol string) group {
	switch symbol {
	case "#":
		return roomOwner()
	case "~":
		return administrator()
	case "&":
		return administrator2()
	case "\u2605":
		return host()
	case "@":
		return moderator()
	case "%":
		return driver()
	case "*":
		return bot()
	case "\u2606":
		return player()
	case "+":
		return voice()
	case " ":
		return normal()
	case "!":
		return muted()
	case "\u2716":
		return namelocked()
	case "\u203d":
		return locked()
	default:
		return normal()
	}
}
