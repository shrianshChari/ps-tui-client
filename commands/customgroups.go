package commands

import (
	"charm-psclient/datastructs"
	"encoding/json"
)

func CustomGroups(data string) (map[string]datastructs.Group, error) {
	var customGroupsList []datastructs.Group
	var customGroups = map[string]datastructs.Group{}

	err := json.Unmarshal([]byte(data), &customGroupsList)
	if err != nil {
		return customGroups, err
	}

	for index, customGroup := range customGroupsList {
		customGroup.Order = index
		customGroups[customGroup.Symbol] = customGroup
	}

	return customGroups, nil
}
