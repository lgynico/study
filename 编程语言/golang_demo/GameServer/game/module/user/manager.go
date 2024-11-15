package user

import "sync"

type group struct {
	sync.Map
}

var groupInstance = &group{}

func (m *group) add(data *UserData) {
	if data == nil {
		return
	}

	m.Store(data.UserID, data)
}

func (m *group) remove(userID int64) {
	m.Delete(userID)
}

func AddToGroup(data *UserData) {
	groupInstance.add(data)
}

func RemoveFromGroup(userID int64) {
	groupInstance.remove(userID)
}

func GetUser(userID int64) (*UserData, bool) {
	value, ok := groupInstance.Load(userID)
	if !ok {
		return nil, false
	}

	return value.(*UserData), true
}

func GetUsers() []*UserData {
	users := make([]*UserData, 0, 16)
	groupInstance.Range(func(key, value any) bool {
		users = append(users, value.(*UserData))
		return true
	})

	return users
}
