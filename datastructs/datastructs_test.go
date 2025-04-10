package datastructs

import (
	"testing"
)

func TestUserLess(t *testing.T) {
	users := UsersSortable{
		User{
			Username: "A",
			Id:       "a",
			Group:    DefaultGroups[" "],
			Away:     false,
		},
		User{
			Username: "B",
			Id:       "b",
			Group:    DefaultGroups["+"],
			Away:     true,
		},
		User{
			Username: "B1",
			Id:       "b1",
			Group:    DefaultGroups["+"],
			Away:     true,
		},
		User{
			Username: "C",
			Id:       "c",
			Group:    DefaultGroups["+"],
			Away:     false,
		},
		User{
			Username: "C1",
			Id:       "c1",
			Group:    DefaultGroups["+"],
			Away:     false,
		},
		User{
			Username: "D",
			Id:       "d",
			Group:    DefaultGroups["~"],
			Away:     true,
		},
		User{
			Username: "E",
			Id:       "e",
			Group:    DefaultGroups["\u203d"],
			Away:     false,
		},
	}

	// Users with a higher rank should be less than users with a lower rank
	if !users.Less(1, 0) {
		// Voiced < Default
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[1].Group.Symbol, users[0].Group.Symbol)
	}
	if !users.Less(5, 1) {
		// Administrator < Voiced
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[5].Group.Symbol, users[1].Group.Symbol)
	}
	if !users.Less(0, 6) {
		// Default < Locked
		t.Fatalf("User with rank '%s' is not less than user with rank '%s'.",
			users[0].Group.Symbol, users[6].Group.Symbol)
	}

	// Non-away users should be less than away users with the same rank
	if users.Less(1, 3) {
		t.Fatalf("Away user is less than non-away user.")
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
