package datastructs

type User struct {
	Username string
	Id       string
	Group    group
	Away     bool
}

type UsersSortable []User

func (u UsersSortable) Len() int      { return len(u) }
func (u UsersSortable) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u UsersSortable) Less(i, j int) bool {
	if u[i].Group.Order != u[j].Group.Order {
		return u[i].Group.Order < u[j].Group.Order
	}
	if !u[i].Away && u[j].Away {
		return false
	}
	if u[i].Away && !u[j].Away {
		return true
	}
	return u[i].Id < u[j].Id
}
