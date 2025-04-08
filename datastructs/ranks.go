package datastructs

type Group struct {
	Name   string
	Symbol string
	Type   string
	Order  int
}

// Group definitions defined in:
// https://github.com/smogon/pokemon-showdown-client/blob/master/play.pokemonshowdown.com/src/client-main.ts#L342
var DefaultGroups = map[string]Group{
	"#": Group{
		Name:   "Room Owner",
		Symbol: "#",
		Type:   "leadership",
		Order:  101,
	},
	"~": Group{
		Name:   "Administrator",
		Symbol: "~",
		Type:   "leadership",
		Order:  102,
	},
	"&": Group{
		Name:   "Administrator",
		Symbol: "&",
		Type:   "leadership",
		Order:  103,
	},
	"\u2605": Group{
		Name:   "Host",
		Symbol: "\u2605",
		Type:   "staff",
		Order:  104,
	},
	"@": Group{
		Name:   "Moderator",
		Symbol: "@",
		Type:   "staff",
		Order:  105,
	},
	"%": Group{
		Name:   "Driver",
		Symbol: "%",
		Type:   "staff",
		Order:  106,
	},
	"*": Group{
		Name:   "Bot",
		Symbol: "*",
		Type:   "normal",
		Order:  109,
	},
	"\u2606": Group{
		Name:   "Player",
		Symbol: "\u2606",
		Type:   "normal",
		Order:  110,
	},
	"+": Group{
		Name:   "Voice",
		Symbol: "+",
		Type:   "normal",
		Order:  200,
	},
	" ": Group{
		Name:   "",
		Symbol: " ",
		Type:   "normal",
		Order:  201,
	},
	"!": Group{
		Name:   "Muted",
		Symbol: "!",
		Type:   "punishment",
		Order:  301,
	},
	"\u2716": Group{
		Name:   "Namelocked",
		Symbol: "\u2716",
		Type:   "punishment",
		Order:  302,
	},
	"\u203d": Group{
		Name:   "Locked",
		Symbol: "\u203d",
		Type:   "punishment",
		Order:  303,
	},
}
