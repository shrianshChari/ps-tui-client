package datastructs

type Room struct {
	RoomName     string
	Users        map[string]User
	ChatMessages ChatMessagesSortable
}

func (r Room) GetUsers() UsersSortable {
	users := UsersSortable{}

	for _, user := range r.Users {
		users = append(users, user)
	}

	return users
}
