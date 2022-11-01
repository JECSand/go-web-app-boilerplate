package models

import "fmt"

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
	uList := NewLinkedList(users, "/admin/", true, true, false, "")
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
	// Init group and user select filter drop down
	var filters []*List
	var overviewScripts []*Script
	//userSelect := NewSelectInput("Select Users", "Select Users", "update", "user_id", "text", GetDataSelectOptions(groupUsers.Users), false)
	userSelect := NewList(groupUsers.Users, "dropdown", "/users", true, "updateTasks")
	filters = append(filters, userSelect)
	if len(groups) > 1 {
		//groupSelect := NewSelectInput("Select Groups", "Select Groups", "update", "group_id", "text", GetDataSelectOptions(groups), false)
		groupSelect := NewList(groups, "dropdown", "/groups", true, "updateTasks")
		filters = append(filters, groupSelect)
	}
	for _, t := range inTasks {
		fmt.Println(t)
	}
	nsTasks, ipTasks, comTasks := SplitTasksByStatus(inTasks)
	// Init Not Started List DIV
	overviewScripts = append(overviewScripts, &Script{Category: "postCheck"})
	notStartedList := NewLinkedList(nsTasks, "/", true, true, true, "postCheck")
	//notStartedList.Script = &Script{Category: "postCheck"}
	notStartedCol := NewListDiv("even columnOne", "", "", NewColumnHeading("Not Started", ""), notStartedList)
	// Init In Progress List DIV
	inProgressList := NewLinkedList(ipTasks, "/", true, true, true, "postCheck")
	//inProgressList.Script = &Script{Category: "postCheck"}
	inProgressCol := NewListDiv("even columnTwo", "", "", NewColumnHeading("In Progress", ""), inProgressList)
	// Init In Completed List DIV
	completedList := NewLinkedList(comTasks, "/", true, true, true, "postCheck")
	//completedList.Script = &Script{Category: "postCheck"}
	completedCol := NewListDiv("even columnThree", "", "", NewColumnHeading("Completed", ""), completedList)
	//	6) Init "drag and drop" jquery script
	//	7) Init and return NewTaskOverview
	return NewTasksOverview("tasksOverview", "", filters, notStartedCol, inProgressCol, completedCol, overviewScripts)
}
