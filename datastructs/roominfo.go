package datastructs

type RoomInfo struct {
	Title     string
	Desc      string
	UserCount int
	Privacy   string
	Section   string
	SubRooms  []string
}

type RoomResponse struct {
	Chat          []RoomInfo
	SectionTitles []string
	UserCount     int
	BattleCount   int
}
