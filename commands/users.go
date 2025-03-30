package commands

import (
	"charm-psclient/datastructs"
	"charm-psclient/utils"
	"sort"
	"strings"
)

func Users(data string) datastructs.UsersSortable {
	var users datastructs.UsersSortable = datastructs.UsersSortable{}

	usersSplit := strings.Split(data, ",")
	for index, user := range usersSplit {
		if index > 0 {
			users = append(users, StringToUser(user))
		}
	}
	sort.Sort(users)
	return users
}

func StringToUser(data string) datastructs.User {
	away := false

	// Necessary to capture groups that are represented by Unicode characters
	runes := []rune(data)

	groupStr := runes[0]
	username := string(runes[1:])

	if strings.HasSuffix(username, "@!") {
		away = true
		username = strings.TrimSuffix(username, "@!")
	}

	return datastructs.User{
		Username: username,
		Id:       utils.ToID(username),
		Group:    datastructs.GetGroup(string(groupStr)),
		Away:     away,
	}
}
