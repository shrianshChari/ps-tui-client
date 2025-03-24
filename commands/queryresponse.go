package commands

import "encoding/json"

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

func QueryresponseRooms(data string) (RoomResponse, error) {
	r := RoomResponse{}

	err := json.Unmarshal([]byte(data), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
