package commands

import (
	"charm-psclient/datastructs"
	"charm-psclient/utils"
	"sort"
	"strings"
)

func Users(data string, groups map[string]datastructs.Group) datastructs.UsersSortable {
	var users datastructs.UsersSortable = datastructs.UsersSortable{}

	usersSplit := strings.Split(data, ",")
	for index, user := range usersSplit {
		if index > 0 {
			users = append(users, StringToUser(user, groups))
		}
	}
	sort.Sort(users)
	return users
}

func StringToUser(data string, groups map[string]datastructs.Group) datastructs.User {
	away := false

	// Necessary to capture groups that are represented by Unicode characters
	runes := []rune(data)

	groupStr := runes[0]
	username := string(runes[1:])

	if strings.HasSuffix(username, "@!") {
		away = true
		username = strings.TrimSuffix(username, "@!")
	}

	group, ok := groups[string(groupStr)]
	if !ok {
		group = groups[" "]
	}

	return datastructs.User{
		Username: username,
		Id:       utils.ToID(username),
		Group:    group,
		Away:     away,
	}
}
