package datastructs

import (
	"testing"
)

func TestUserLess(t *testing.T) {
	users := UsersSortable{
		User{
			Username: "A",
			Id:       "a",
			Group:    GetGroup(" "),
			Away:     false,
		},
		User{
			Username: "B",
			Id:       "b",
			Group:    GetGroup("+"),
			Away:     true,
		},
		User{
			Username: "B1",
			Id:       "b1",
			Group:    GetGroup("+"),
			Away:     true,
		},
		User{
			Username: "C",
			Id:       "c",
			Group:    GetGroup("+"),
			Away:     false,
		},
		User{
			Username: "C1",
			Id:       "c1",
			Group:    GetGroup("+"),
			Away:     false,
		},
		User{
			Username: "D",
			Id:       "d",
			Group:    GetGroup("~"),
			Away:     true,
		},
		User{
			Username: "E",
			Id:       "e",
			Group:    GetGroup("\u203d"),
			Away:     false,
		},
	}

	// Users with a lower rank should be less than users with a higher rank
	if !users.Less(0, 1) {
		// Default vs Voiced
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[0].Group.Symbol, users[1].Group.Symbol)
	}
	if !users.Less(1, 5) {
		// Voiced vs Administrator
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[1].Group.Symbol, users[5].Group.Symbol)
	}
	if !users.Less(6, 0) {
		// Locked vs Default
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[6].Group.Symbol, users[0].Group.Symbol)
	}

	// Away users should be less than non-away users with the same rank
	if !users.Less(1, 3) {
		t.Fatalf("Away user is not less than non-away user.")
	}

	// Users with the same rank and away status should be sorted by ID
	if !users.Less(1, 2) {
		t.Fatalf("User with ID '%s' is not less than user with ID '%s'.",
			users[1].Id, users[2].Id)
	}
	if !users.Less(3, 4) {
		t.Fatalf("User with ID '%s' is not less than user with ID '%s'.",
			users[3].Id, users[4].Id)
	}
}
