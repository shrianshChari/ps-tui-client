package datastructs

type Server struct {
	Rooms     map[string]Room
	RoomsInfo RoomResponse
	Groups    map[string]Group
}
