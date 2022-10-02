package models

// InitializeUserSettings ...
func InitializeUserSettings(user *User) *Settings {
	SettingsForm := InitializeUserSettingsForm(user)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	infoLink := NewLink("active", "", "/account/settings", "Change User Info", true)
	pwLink := NewLink("", "", "/account/settings/password", "Change Password", true)
	optionsCol := NewLinkDiv("columnOne", "", "", NewColumnHeading("Options", ""), []*Link{infoLink, pwLink})
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	unLink := NewLink("active", "", "", user.Username, true)
	rLink := NewLink("active", "", "", user.Role, true)
	infoCol := NewLinkDiv("columnTwo", "", "", NewColumnHeading("Current Info", ""), []*Link{unLink, rLink})
	return NewSettings("", "", SettingsForm, optionsCol, infoCol)
}

// InitializeGroupSettings ...
func InitializeGroupSettings(group *Group, users []*User) *Settings {
	SettingsForm := InitializeGroupSettingsForm(group)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	uList := NewLinkedUsersList(users)
	usersCol := NewListDiv("columnOne", "", "", NewColumnHeading("Group Users", ""), uList)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	nameLink := NewLink("active", "", "", group.Name, true)
	infoCol := NewLinkDiv("columnTwo", "", "", NewColumnHeading("Group Info", ""), []*Link{nameLink})
	return NewSettings("", "", SettingsForm, usersCol, infoCol)
}
