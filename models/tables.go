package models

// UserTable ...
type UserTable struct {
	CreateUserForm CreateUserForm
	UpdateUserForm UpdateUserForm
	Users          []*User
}

// Initialize ...
func (ut *UserTable) Initialize(users []*User, method string, id string) {
	if method == "GET" {
		if id == "" {
			ut.Users = users
		} else if id == "create" {
			ut.CreateUserForm = CreateUserForm{}
		} else {
			upUserForm := UpdateUserForm{}
			upUserForm.Load(users[0])
			ut.UpdateUserForm = upUserForm
		}
	}
}

/*
// findUserByUuid
func (ut *UserTable) findUserByUuid(userUuid string) User {
	var user User
	for _, u := range ut.Users {
		if u.Uuid == userUuid {
			user = u
			break
		}
	}
	return user
}

// Load
func (ut *UserTable) Load(auth Auth) {
	var users []User
	var user User
	users = user.GetAll(auth)
	ut.Users = users
}

// LoadForm
func (ut *UserTable) LoadForm(userUuid string) {
	user := ut.findUserByUuid(userUuid)
	if userUuid != "" {
		// Create User
		ut.UpdateUserForm.Load(user)
	}
}
*/

/*
////////////////////////////////////////////
*/

// GroupTable ...
type GroupTable struct {
	CreateGroupForm CreateGroupForm
	UpdateGroupForm UpdateGroupForm
	Groups          []*Group
}

// Initialize ...
func (gt *GroupTable) Initialize(groups []*Group, method string, id string) {
	if method == "GET" {
		if id == "" {
			gt.Groups = groups
		} else if id == "create" {
			gt.CreateGroupForm = CreateGroupForm{}
		} else {
			upGroupForm := UpdateGroupForm{}
			upGroupForm.Load(groups[0])
			gt.UpdateGroupForm = upGroupForm
		}
	}
}

/*
// Load
func (gt *GroupTable) findGroupByUuid(groupUuid string) Group {
	var group Group
	for _, gt := range gt.Groups {
		if gt.Uuid == groupUuid {
			group = gt
			break
		}
	}
	return group
}

// Load
func (gt *GroupTable) Load(auth Auth) {
	var groups []Group
	var group Group
	groups = group.GetAll(auth)
	gt.Groups = groups
}

// LoadForm
func (gt *GroupTable) LoadForm(groupUuid string) {
	group := gt.findGroupByUuid(groupUuid)
	if groupUuid != "" {
		// Create User
		gt.UpdateGroupForm.Load(group)
	}
}
*/
