package commands

import (
	"charm-psclient/datastructs"
	"encoding/json"
)

func QueryresponseRooms(data string) (datastructs.RoomResponse, error) {
	r := datastructs.RoomResponse{}

	err := json.Unmarshal([]byte(data), &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
