package models

// InitializeUserSettings ...
func InitializeUserSettings(user *User, admin bool) *Settings {
	SettingsForm := InitializeUserSettingsForm(user, admin)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	var links []*Link
	infoLink := NewLink("active", "", "/account/settings", "Change User Info", true)
	links = append(links, infoLink)
	if !admin {
		pwLink := NewLink("", "", "/account/password", "Change Password", true)
		links = append(links, pwLink)
	}
	optionsCol := NewLinkDiv("uneven columnOne", "", "", NewColumnHeading("Options", ""), links)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	unLink := NewLink("active", "", "", user.Username, true)
	rLink := NewLink("active", "", "", user.Role, true)
	infoCol := NewLinkDiv("uneven columnTwo", "", "", NewColumnHeading("Current Info", ""), []*Link{unLink, rLink})
	return NewSettings("", "", SettingsForm, optionsCol, infoCol)
}

// InitializeGroupSettings ...
func InitializeGroupSettings(group *Group, users []*User) *Settings {
	SettingsForm := InitializeGroupSettingsForm(group)
	// 1. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	uList := NewLinkedList(users, "/admin/", true, true, false)
	usersCol := NewListDiv("uneven columnOne", "", "", NewColumnHeading("Group Users", ""), uList)
	// 2. NewLinkDiv(class string, id string, label string, head *Heading, links []*Link)
	nameLink := NewLink("active", "", "", group.Name, true)
	infoCol := NewLinkDiv("uneven columnTwo", "", "", NewColumnHeading("Group Info", ""), []*Link{nameLink})
	return NewSettings("", "", SettingsForm, usersCol, infoCol)
}

/*
NOTES:
	*EACH TASK LIST WILL HAVE TWO AXIOS CALLS:
		1) POST UpdateTaskStatus based on drag and drop
		2) GET	Sporadic GetTasks to get current task updates for keeping sync

	*THESE AXIOS REQUESTS HIT A FRONTEND ROUTE WHICH PROXIES TO THE BACKEND API FOR DATA

	*INITIAL GROUP TASKS ARE PROVIDED BY CONTROLLER AND LOADED INTO PAGE

	*FILTER SELECT BUTTONS:
		1) USERS: FILTERS LIST ITEMS USING JQUERY
		2) GROUPS: RELOADS PAGE FOR A GIVEN GROUP (RootAdmin only)
*/

// InitializeTaskOverview instantiates a task Overview Abstract	// TODO NEXT - START HERE, called in TASK CONTROLLER
func InitializeTaskOverview(inTasks []*Task, groupUsers *GroupUsersDTO, groups []*Group) *Overview {
	nsTasks, ipTasks, comTasks := SplitTasksByStatus(inTasks)
	// STEPS:
	//	1) Init Not Started List DIV
	notStartedList := NewLinkedList(nsTasks, "/", true, true, true)
	notStartedList.Script = &Script{Category: "postCheck"}
	notStartedCol := NewListDiv("even columnOne", "", "", NewColumnHeading("Not Started", ""), notStartedList)
	//	2) Init In Progress List DIV
	inProgressList := NewLinkedList(ipTasks, "/", true, true, true)
	inProgressList.Script = &Script{Category: "postCheck"}
	inProgressCol := NewListDiv("even columnTwo", "", "", NewColumnHeading("In Progress", ""), inProgressList)
	//  3) Init In Completed List DIV
	completedList := NewLinkedList(comTasks, "/", true, true, true)
	completedList.Script = &Script{Category: "postCheck"}
	completedCol := NewListDiv("even columnThree", "", "", NewColumnHeading("Completed", ""), completedList)
	//	4) Init group select filter drop down (master admin)
	// 	5) Init user select filter drop down
	//	6) Init "drag and drop" jquery script
	//	7) Init and return NewTaskOverview
	return NewTasksOverview("tasksOverview", "", notStartedCol, inProgressCol, completedCol)
}
