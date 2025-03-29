package datastructs

type Room struct {
	RoomName     string
	Users        map[string]User
	ChatMessages ChatMessagesSortable
}
